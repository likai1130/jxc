
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="客户名称" prop="customerName"><el-input v-model="queryParams.customerName" placeholder="请输入客户名称" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="客户类型" prop="customerType"><el-select v-model="queryParams.customerType"
                                               placeholder="客户信息管理客户类型" clearable size="small">
                                        <el-option
                                                v-for="dict in customerTypeOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['jxc:consumer:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['jxc:consumer:edit']"
                                type="success"
                                icon="el-icon-edit"
                                size="mini"
                                :disabled="single"
                                @click="handleUpdate"
                        >修改
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['jxc:consumer:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="consumerList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="客户名称" align="center" prop="customerName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="客户类型" align="center" prop="customerType"
                                                 :formatter="customerTypeFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ customerTypeFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="联系人姓名" align="center" prop="contactName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="联系电话" align="center" prop="contactPhone"
                                                 :show-overflow-tooltip="true"/><el-table-column label="电子邮箱" align="center" prop="email"
                                                 :show-overflow-tooltip="true"/><el-table-column label="公司地址" align="center" prop="companyAddress"
                                                 :show-overflow-tooltip="true"/><el-table-column label="送货地址" align="center" prop="deliveryAddress"
                                                 :show-overflow-tooltip="true"/><el-table-column label="备注信息" align="center" prop="notes"
                                                 :show-overflow-tooltip="true"/>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                         <el-popconfirm
                           class="delete-popconfirm"
                           title="确认要修改吗?"
                           confirm-button-text="修改"
                           @confirm="handleUpdate(scope.row)"
                         >
                           <el-button
                             slot="reference"
                             v-permisaction="['jxc:consumer:edit']"
                             size="mini"
                             type="text"
                             icon="el-icon-edit"
                           >修改
                           </el-button>
                         </el-popconfirm>
                         <el-popconfirm
                            class="delete-popconfirm"
                            title="确认要删除吗?"
                            confirm-button-text="删除"
                            @confirm="handleDelete(scope.row)"
                         >
                            <el-button
                              slot="reference"
                              v-permisaction="['jxc:consumer:remove']"
                              size="mini"
                              type="text"
                              icon="el-icon-delete"
                            >删除
                            </el-button>
                         </el-popconfirm>
                        </template>
                    </el-table-column>
                </el-table>

                <pagination
                        v-show="total>0"
                        :total="total"
                        :page.sync="queryParams.pageIndex"
                        :limit.sync="queryParams.pageSize"
                        @pagination="getList"
                />

                <!-- 添加或修改对话框 -->
                <el-dialog :title="title" :visible.sync="open" width="500px">
                    <el-form ref="form" :model="form" :rules="rules" label-width="100px">
                        
                                    <el-form-item label="客户名称" prop="customerName">
                                        <el-input v-model="form.customerName" placeholder="客户名称"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="客户类型" prop="customerType">
                                        <el-select v-model="form.customerType"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in customerTypeOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="联系人姓名" prop="contactName">
                                        <el-input v-model="form.contactName" placeholder="联系人姓名"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="联系电话" prop="contactPhone">
                                        <el-input v-model="form.contactPhone" placeholder="联系电话"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="电子邮箱" prop="email">
                                        <el-input v-model="form.email" placeholder="电子邮箱"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="公司地址" prop="companyAddress">
                                        <el-input v-model="form.companyAddress" placeholder="公司地址"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="送货地址" prop="deliveryAddress">
                                        <el-input v-model="form.deliveryAddress" placeholder="送货地址"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="备注信息" prop="notes">
                                        <el-input v-model="form.notes" placeholder="备注信息"
                                                      />
                                    </el-form-item>
                    </el-form>
                    <div slot="footer" class="dialog-footer">
                        <el-button type="primary" @click="submitForm">确 定</el-button>
                        <el-button @click="cancel">取 消</el-button>
                    </div>
                </el-dialog>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {addConsumer, delConsumer, getConsumer, listConsumer, updateConsumer} from '@/api/jxc/tb-consumer'
    
    export default {
        name: 'Consumer',
        components: {
        },
        data() {
            return {
                // 遮罩层
                loading: true,
                // 选中数组
                ids: [],
                // 非单个禁用
                single: true,
                // 非多个禁用
                multiple: true,
                // 总条数
                total: 0,
                // 弹出层标题
                title: '',
                // 是否显示弹出层
                open: false,
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                consumerList: [],
                customerTypeOptions: [],
                // 关系表类型
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    customerName:undefined,
                    customerType:undefined,
                    createdAtOrder: "desc"

                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {customerName:  [ {required: true, message: '客户名称不能为空', trigger: 'blur'} ],
                customerType:  [ {required: true, message: '客户类型不能为空', trigger: 'blur'} ],
                contactName:  [ {required: true, message: '联系人名称不能为空', trigger: 'blur'} ],
                contactPhone:  [ {required: true, message: '联系人电话不能为空', trigger: 'blur'} ],
                deliveryAddress:  [ {required: true, message: '送货地址不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getDicts('bus_customer_type').then(response => {
                this.customerTypeOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listConsumer(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.consumerList = response.data.list
                        this.total = response.data.count
                        this.loading = false
                    }
                )
            },
            // 取消按钮
            cancel() {
                this.open = false
                this.reset()
            },
            // 表单重置
            reset() {
                this.form = {
                
                id: undefined,
                customerName: undefined,
                customerType: undefined,
                contactName: undefined,
                contactPhone: undefined,
                email: undefined,
                companyAddress: undefined,
                deliveryAddress: undefined,
                notes: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            customerTypeFormat(row) {
                return this.selectDictLabel(this.customerTypeOptions, row.customerType)
            },
            // 关系
            // 文件
            /** 搜索按钮操作 */
            handleQuery() {
                this.queryParams.pageIndex = 1
                this.getList()
            },
            /** 重置按钮操作 */
            resetQuery() {
                this.dateRange = []
                this.resetForm('queryForm')
                this.handleQuery()
            },
            /** 新增按钮操作 */
            handleAdd() {
                this.reset()
                this.open = true
                this.title = '添加客户信息管理'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.id)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 修改按钮操作 */
            handleUpdate(row) {
                this.reset()
                const id =
                row.id || this.ids
                getConsumer(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改客户信息管理'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateConsumer(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addConsumer(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        }
                    }
                })
            },
            /** 删除按钮操作 */
            handleDelete(row) {
                var Ids = (row.id && [row.id]) || this.ids

                this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(function () {
                      return delConsumer( { 'ids': Ids })
                }).then((response) => {
                   if (response.code === 200) {
                     this.msgSuccess(response.msg)
                     this.open = false
                     this.getList()
                   } else {
                     this.msgError(response.msg)
                   }
                }).catch(function () {
                })
            }
        }
    }
</script>
