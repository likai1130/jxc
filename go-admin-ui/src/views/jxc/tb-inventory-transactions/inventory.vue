
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
            <h2>商品入库</h2>
                <el-form ref="form" :inline="true" :model="form" :rules="rules" label-width="120px">
                    <el-form-item label="采购单号" prop="orderNumber"> 
                        <el-select v-model="form.orderNumber" clearable 
                                         @change="getPurchaseListGoods"
                                         placeholder="请选择" >
                                <el-option
                                        v-for="dict in purchaseNumberOptions"
                                        :key="dict.key"
                                        :label="dict.value"
                                        :value="dict.key"
                                />
                            </el-select>
                    </el-form-item>
                    <el-form-item label="备注" prop="remarks">
                        <el-input v-model="form.remarks" placeholder="备注"
                                        />
                    </el-form-item>
                    <el-form-item label="库存类型" prop="remarks"  v-if="false">
                        <el-input v-model="form.transactionType" placeholder="库存类型"
                                        />
                    </el-form-item>
                    </br>
                    <h3>相关商品列表</h3>
                    <el-table v-loading="goodsLoading" :data="purchaseListGoodsList">
                    <el-table-column  width="55" align="center"/><el-table-column label="商品编码" align="center" prop="gCode"
                        :show-overflow-tooltip="true"/><el-table-column label="商品名称" align="center" prop="gName"
                        :show-overflow-tooltip="true"/><el-table-column label="商品规格" align="center" prop="specification"
                        :show-overflow-tooltip="true"/><el-table-column label="单位" align="center" prop="unit"
                        :show-overflow-tooltip="true"/><el-table-column label="单价(￥)" align="center" prop="price"
                        :show-overflow-tooltip="true"/><el-table-column label="数量" align="center" prop="num"
                        :show-overflow-tooltip="true"/><el-table-column label="商品" v-if="false" align="center" prop="goodsId"
                        :show-overflow-tooltip="true"/>
                </el-table>
                    <el-form-item size="small">
                        <el-button type="primary" @click="submitForm">入库</el-button>
                        <el-button @click="reset">重置</el-button>
                        </el-form-item>
                </el-form>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {listPurchaseList} from '@/api/jxc/tb-purchase-list'
    import {listPurchaseListGoods} from '@/api/jxc/tb-purchase-list-goods'
    import {addInventoryTransactions,updateInventoryTransactions} from '@/api/jxc/tb-inventory-transactions'

    
    export default {
        name: 'PurchaseList',
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
                // 类型数据字典
                typeOptions: [],
                purchaseListList: [],
                storageStatusOptions: [],
                // 关系表类型
                purchaseNumberOptions: [],
                purchaseListGoodsList: [],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10000,
                    storageStatus:"0",
                    createdAtOrder: "desc"    
                    
                },
                // 表单参数
                form: {
                },
                goodsLoading: false,
                goodsQueryParams:{
                    pageIndex: 1,
                    pageSize: 10000,
                    purchaseListId: undefined,
                },
                // 表单校验
                 rules: {
                    orderNumber:  [ {required: true, message: '采购订单不能为空', trigger: 'blur'} ]
                },
        }
        },
        created() {
            this.getPurchaseListOptionsItems()
        },
        methods: {
            // 采购单关系
            purchaseNumberFormat(row) {
                return this.selectItemsLabel(this.purchaseNumberOptions, row.purchaseNumber)
            },
             // 关系
             getPurchaseListOptionsItems() {
               this.getItems(listPurchaseList, this.queryParams).then(res => {
                   this.purchaseNumberOptions = this.setItems(res, 'id', 'purchaseNumber')
               })
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
                amountPayable: undefined,
                amountPaid: undefined,
                purchaseDate: undefined,
                remarks: undefined,
                customerId: undefined,
                purchaseNumber:undefined,
                orderNumber: undefined,

            }
                this.resetForm('form')
                this.purchaseListGoodsList = []
                this.getPurchaseListOptionsItems()
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
           
            /** 重置按钮操作 */
            resetQuery() {
                this.dateRange = []
                this.resetForm('queryForm')
                this.handleQuery()
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.id)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
          
            /**销售商品明细列表 */
            getPurchaseListGoods(id) {
                this.goodsQueryParams.purchaseListId = id
                this.goodsLoading = true
                listPurchaseListGoods(this.goodsQueryParams).then(response => {
                        this.purchaseListGoodsList = response.data.list
                        this.goodsLoading = false
                    }
                )
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        this.form.transactionType = "0"
                        const selectedOrder = this.purchaseNumberOptions.find(option => option.key === this.form.orderNumber);
                        const payload = {
                            ...this.form,
                            orderId: selectedOrder ? Number(selectedOrder.key) : 0, // 订单 ID 转换为整数
                            orderNumber: selectedOrder ? selectedOrder.value : '', // 使用订单单号
                            goods: this.purchaseListGoodsList // 将商品数组作为表单的一部分
                        };
                        addInventoryTransactions(payload).then(response => {
                            if (response.code === 200) {
                                this.msgSuccess(response.msg)
                                this.reset()
                            } else {
                                this.msgError(response.msg)
                            }
                        })
                    }
                })
            }
        }
    }
</script>
