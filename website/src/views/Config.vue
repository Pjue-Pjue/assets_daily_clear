<template>
    <section>
        <div>
            <el-button @click="addDialogVisible=true">Add Config</el-button>
            <el-button @click="StartWork">Start Work</el-button>
            <el-button @click="StopWork">Stop Work</el-button>
            <el-button @click="DeleteItems">Delete</el-button>
        </div>
        <div style="margin-top: 10px;margin-left: 5%; width: 90%;text-align: center">
            <el-table :data="datalist" @selection-change="handleSelection" :row-class-name="tableRowClassName">
                <el-table-column type="selection" width="55">
                </el-table-column>
                <el-table-column label="Market" prop="Market" width="100"></el-table-column>
                <el-table-column label="Api Key" prop="AccessKey"></el-table-column>
                <el-table-column label="Secret Key" prop="SecretKey">
                </el-table-column>
                <el-table-column label="PassPhrase" prop="PassPhrase">
                </el-table-column>
                <el-table-column label="Label" prop="Label" width="120">
                </el-table-column>
                <el-table-column label="Status" prop="Status" width="80">
                    <template slot-scope="scope">
                        {{StatusEnum[scope.row.Status]}}
                    </template>
                </el-table-column>
                <el-table-column label="Operation">
                    <template slot-scope="scope">
<!--                        <el-button @click="EditItem(scope.row.Id)">Edit</el-button>-->
<!--                        <el-button @click="DeleteItem(scope.row.Id)">Delete</el-button>-->
                        <el-button @click="ShowDetails(scope.row.Id)">Query</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <el-dialog title="Add Account" :visible.sync="addDialogVisible" width="500px" center custom-class="success-modal" :show-close="false" :close-on-click-modal="false">
            <el-form :model="AddForm" :rules="addRules" ref="AddForm" :label-width="formLabelWidth">
                <el-form-item label="market" prop="Market">
                    <el-select v-model="AddForm.Market" placeholder="choose market" style="width: 85%">
                        <el-option v-for="item in marketItems" :label="item.Name" :value="item.Id" :key="item.Id">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Api Key" prop="ApiKey">
                    <el-input v-model="AddForm.ApiKey" autocomplete="off" style="width: 85%"></el-input>
                </el-form-item>
                <el-form-item label="Secret Key" prop="SecretKey">
                    <el-input v-model="AddForm.SecretKey" autocomplete="off" style="width: 85%"></el-input>
                </el-form-item>
                <el-form-item label="PassPhrase">
                    <el-input v-model="AddForm.PassPhrase" autocomplete="off" style="width: 85%"></el-input>
                </el-form-item>
                <el-form-item label="Label">
                    <el-input v-model="AddForm.Label" autocomplete="off" style="width: 85%"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button type="primary" @click="submitAddForm('AddForm')">confirm</el-button>
                <el-button @click="resetFormFields('AddForm')">cancel</el-button>
            </div>
        </el-dialog>
        <el-dialog :title="OperationTitle" :visible.sync="OperationDialogVisible" width="30%" center
                   custom-class="danger-modal" :show-close="false">
            <span style="display: block; text-align: center; line-height: 60px; font-size: 15px;">Are You Sure to <span style="color: red;">{{UpdateText}}</span> ？</span>
            <span slot="footer" class="dialog-footer">
        <el-button @click="Confirm">Sure</el-button>
        <el-button @click="OperationDialogVisible = false">Cancel</el-button>
      </span>
        </el-dialog>
    </section>
</template>
<script>
    import axios from 'axios'
    export default {
        data () {
           return {
               OperationTitle: '',
               UpdateIds: [],
               OperationDialogVisible: false,
               UpdateText: '',
               OperationFlag: '',
               UpdateStatusEnum: {
                   '0': 0, // delete
                   '1': 2, // start
                   '2': 1, // stop
               },
               rowChecked: [],
               multipleSelection: [],
               marketItems:[
                   {
                       Id: 'HUOBI',
                       Name: 'HUOBI'
                   }, {
                       Id: 'OKEX',
                       Name: 'OKEX'
                   },
               ],
               formInputWidth: '200px',
               formLabelWidth: '120px',
               AddForm: {
                   Market: '',
                   ApiKey: '',
                   SecretKey: '',
                   PassPhrase: '',
                   Label: ''
               },
               addRules: {
                   Market: [
                       {required: true, message: 'please choose market !!!'}
                   ],
                   ApiKey: [
                       { required: true, message: 'ApiKey is required', trigger: 'change' }
                   ],
                   SecretKey: [
                       { required: true, message: 'SecretKey is required', trigger: 'change' }
                   ]
               },
               addDialogVisible: false,
               StatusEnum: {
                   0: 'Deleted',
                   1: 'Stopped',
                   2: 'Running'
               },
                datalist: []
            }
        },
        created () {
            axios.get('/api/getAccounts').then(response => {
                let code = response.data.code
                if (code === 200) {
                    this.datalist = response.data.data
                } else {
                    let msg = response.data.msg
                    this.$message.error(msg)
                }
            })
        },
        methods: {
            submitAddForm (addForm) {
                this.$refs[addForm].validate((valid) => {
                    if (valid) {
                        axios.post('/api/addAccount', {
                            'market': this.AddForm.Market,
                            'access_key': this.AddForm.ApiKey,
                            'secret_key': this.AddForm.SecretKey,
                            'pass_phrase': this.AddForm.PassPhrase,
                            'label': this.AddForm.Label
                        }).then(response => {
                            let code = response.data.code
                            if (code === 200) {
                                this.$message.success('添加成功')
                                this.$refs[addForm].resetFields()
                                this.addDialogVisible = false
                                axios.get('/api/getAccounts').then(response => {
                                    let code = response.data.code
                                    if (code === 200) {
                                        this.datalist = response.data.data
                                    } else {
                                        let msg = response.data.msg
                                        this.$message.error(msg)
                                    }
                                })
                            } else {
                                let msg = response.data.msg
                                this.$message.error(msg)
                            }
                        })
                    }
                })
            },
            resetFormFields (form) {
                this.addDialogVisible = false
                this.$refs[form].resetFields()
            },
            // EditItem (id) {},
            DeleteItems () {
                this.UpdateIds = []
                if (this.multipleSelection.length === 0) {
                    this.$message.error('Please Choose Item')
                    return
                } else if (this.multipleSelection.length === 1) {
                    if (this.multipleSelection[0].Status !== 1) {
                        this.$message.error('only stopped account can be deleted')
                        return
                    }
                    this.UpdateIds.push(this.multipleSelection[0].Id)
                    this.UpdateText = ' Delete this Account'
                } else {
                    for (let i = 0; i < this.multipleSelection.length; i++) {
                        if (this.multipleSelection[i].Status !== 1) {
                            this.$message.error('only stopped account can be deleted')
                            return
                        }
                    }
                    let ids = []
                    this.multipleSelection.forEach(row => {
                        ids.push(row.Id)
                    })
                    this.UpdateIds = ids
                    this.UpdateText = ' Delete these Accounts'
                }
                this.OperationTitle = 'Delete'
                this.OperationDialogVisible = true
                this.OperationFlag = 0
            },
            StartWork () {
                this.UpdateIds = []
                if (this.multipleSelection.length === 0) {
                    this.$message.error('Please Choose Item')
                    return
                } else if (this.multipleSelection.length === 1) {
                    if (this.multipleSelection[0].Status !== 1) {
                        this.$message.error('only stopped account can do this')
                        return
                    }
                    this.UpdateIds.push(this.multipleSelection[0].Id)
                    this.UpdateText = 'Start Statistical this Account'
                } else {
                    for (let i = 0; i < this.multipleSelection.length; i++) {
                        if (this.multipleSelection[i].Status !== 1) {
                            this.$message.error('only stopped account can do this')
                            return
                        }
                    }
                    let ids = []
                    this.multipleSelection.forEach(row => {
                        ids.push(row.Id)
                    })
                    this.UpdateIds = ids
                    this.UpdateText = 'Start Statistical these Accounts'
                }
                this.OperationTitle = 'Start'
                this.OperationDialogVisible = true
                this.OperationFlag = 1
            },
            StopWork () {
                this.UpdateIds = []
                if (this.multipleSelection.length === 0) {
                    this.$message.error('Please Choose Item')
                    return
                } else if (this.multipleSelection.length === 1) {
                    if (this.multipleSelection[0].Status !== 2) {
                        this.$message.error('only running account can be stopped')
                        return
                    }
                    this.UpdateIds.push(this.multipleSelection[0].Id)
                    this.UpdateText = 'Stop Statistical this Account'
                } else {
                    for (let i = 0; i < this.multipleSelection.length; i++) {
                        if (this.multipleSelection[i].Status !== 2) {
                            this.$message.error('only running account can be stopped')
                            return
                        }
                    }
                    let ids = []
                    this.multipleSelection.forEach(row => {
                        ids.push(row.Id)
                    })
                    this.UpdateIds = ids
                    this.UpdateText = 'Stop Statistical these Accounts'
                }
                this.OperationTitle = 'Stop'
                this.OperationDialogVisible = true
                this.OperationFlag = 2
            },
            Confirm () {
                if (this.OperationFlag === '' || this.OperationFlag === null || this.OperationFlag ===undefined) {
                    this.$message.error('internal error')
                    return
                }
                if (this.UpdateIds.length <= 0) {
                    this.$message.error('no operation item')
                    return
                }
                axios.post('/api/updateAccount', {
                    'ids': this.UpdateIds,
                    'status': this.UpdateStatusEnum[this.OperationFlag]
                }).then(response => {
                    let code = response.data.code
                    if (code === 200) {
                        this.$message.success('操作成功')
                        this.OperationDialogVisible = false
                        axios.get('/api/getAccounts').then(response => {
                            let code = response.data.code
                            if (code === 200) {
                                this.datalist = response.data.data
                            } else {
                                let msg = response.data.msg
                                this.$message.error(msg)
                            }
                        })
                    } else {
                        let msg = response.data.msg
                        this.$message.error(msg)
                    }
                })
            },
            ShowDetails (id) {},
            handleSelection (selection, row) {
                this.rowChecked = []
                this.multipleSelection = selection
                let ids = []
                this.datalist.map((c, i) => {
                    selection.map((s, j) => {
                        if (c.id === s.id) {
                            ids.push(i)
                        }
                    })
                })
                this.rowChecked = ids
            },
            tableRowClassName ({row, rowIndex}) {
                if (this.rowChecked.indexOf(rowIndex) >= 0) {
                    return 'checked-bg'
                } else {
                    return ''
                }
            }
        }
    }
</script>
<style scoped>
    .el-table .checked-bg {
        background: #FFFFCC;
    }
</style>