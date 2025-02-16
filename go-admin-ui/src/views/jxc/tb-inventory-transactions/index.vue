
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="操作类型" prop="transactionType"><el-select v-model="queryParams.transactionType"
                                               placeholder="库存管理操作类型" clearable size="small">
                                        <el-option
                                                v-for="dict in transactionTypeOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="商品编码" prop="gCode"><el-input v-model="queryParams.gCode" placeholder="请输入商品编码" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="商品名称" prop="gName"><el-input v-model="queryParams.gName" placeholder="请输入商品名称" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="订单号" prop="orderNumber"><el-input v-model="queryParams.orderNumber" placeholder="请输入订单号" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <!-- <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['jxc:inventoryTransactions:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['jxc:inventoryTransactions:edit']"
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
                                v-permisaction="['jxc:inventoryTransactions:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row> -->

                <el-table v-loading="loading" :data="inventoryTransactionsList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="订单号" align="center" prop="orderNumber"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品编码" align="center" prop="gCode"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品名称" align="center" prop="gName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品规格" align="center" prop="specification"
                                                 :show-overflow-tooltip="true"/><el-table-column label="数量" align="center" prop="num"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品ID" align="center" v-if="false"  prop="goodsId"
                                                 :show-overflow-tooltip="true"/><el-table-column label="订单ID" align="center"  v-if="false"  prop="orderId"
                                                 :show-overflow-tooltip="true"/>
                                <el-table-column label="操作类型" align="center" prop="transactionType"
                                    :formatter="transactionTypeFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ transactionTypeFormat(scope.row) }}
                                    </template>
                                </el-table-column>
                                <el-table-column label="备注" align="center" prop="remarks"
                                                 :show-overflow-tooltip="true"/>
                                 <el-table-column label="操作时间" align="center" prop="createdAt" width="180">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.createdAt) }}</span>
                                    </template>
                                </el-table-column>
                    <!-- <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                         <el-popconfirm
                           class="delete-popconfirm"
                           title="确认要修改吗?"
                           confirm-button-text="修改"
                           @confirm="handleUpdate(scope.row)"
                         >
                           <el-button
                             slot="reference"
                             v-permisaction="['jxc:inventoryTransactions:edit']"
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
                              v-permisaction="['jxc:inventoryTransactions:remove']"
                              size="mini"
                              type="text"
                              icon="el-icon-delete"
                            >删除
                            </el-button>
                         </el-popconfirm>
                        </template>
                    </el-table-column> -->
                </el-table>

                <pagination
                        v-show="total>0"
                        :total="total"
                        :page.sync="queryParams.pageIndex"
                        :limit.sync="queryParams.pageSize"
                        @pagination="getList"
                />

                <!-- 添加或修改对话框 -->
                <!-- <el-dialog :title="title" :visible.sync="open" width="500px">
                    <el-form ref="form" :model="form" :rules="rules" label-width="80px">
                        
                                    <el-form-item label="操作类型" prop="transactionType">
                                        <el-select v-model="form.transactionType"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in transactionTypeOptions"
                                                        :key="dict.value"
                                                        :label="dict.label"
                                                        :value="dict.value"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="商品编码" prop="gCode">
                                        <el-input v-model="form.gCode" placeholder="商品编码"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品名称" prop="gName">
                                        <el-input v-model="form.gName" placeholder="商品名称"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="数量" prop="num">
                                        <el-input v-model="form.num" placeholder="数量"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单号" prop="orderNumber">
                                        <el-input v-model="form.orderNumber" placeholder="订单号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品ID" prop="goodsId">
                                        <el-input v-model="form.goodsId" placeholder="商品ID"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="订单ID" prop="orderId">
                                        <el-input v-model="form.orderId" placeholder="订单ID"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="备注" prop="remarks">
                                        <el-input v-model="form.remarks" placeholder="备注"
                                                      />
                                    </el-form-item>
                    </el-form>
                    <div slot="footer" class="dialog-footer">
                        <el-button type="primary" @click="submitForm">确 定</el-button>
                        <el-button @click="cancel">取 消</el-button>
                    </div>
                </el-dialog> -->
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {addInventoryTransactions, delInventoryTransactions, getInventoryTransactions, listInventoryTransactions, updateInventoryTransactions} from '@/api/jxc/tb-inventory-transactions'
    
    export default {
        name: 'InventoryTransactions',
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
                inventoryTransactionsList: [],
                transactionTypeOptions: [],
                // 关系表类型
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    transactionType:undefined,
                    gCode:undefined,
                    gName:undefined,
                    orderNumber:undefined,
                    goodsId:undefined,
                    orderId:undefined,
                    createdAtOrder: "desc"
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {transactionType:  [ {required: true, message: '操作类型不能为空', trigger: 'blur'} ],
                gCode:  [ {required: true, message: '商品编码不能为空', trigger: 'blur'} ],
                gName:  [ {required: true, message: '商品名称不能为空', trigger: 'blur'} ],
                orderNumber:  [ {required: true, message: '订单号不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getDicts('bus_inventory_transaction_type').then(response => {
                this.transactionTypeOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listInventoryTransactions(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.inventoryTransactionsList = response.data.list
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
                transactionType: undefined,
                gCode: undefined,
                gName: undefined,
                num: undefined,
                orderNumber: undefined,
                goodsId: undefined,
                orderId: undefined,
                remarks: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            transactionTypeFormat(row) {
                return this.selectDictLabel(this.transactionTypeOptions, row.transactionType)
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
                this.title = '添加库存管理'
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
                getInventoryTransactions(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改库存管理'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateInventoryTransactions(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addInventoryTransactions(this.form).then(response => {
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
                      return delInventoryTransactions( { 'ids': Ids })
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
