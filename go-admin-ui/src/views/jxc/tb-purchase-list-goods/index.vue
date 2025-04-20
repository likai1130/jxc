
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="form" :inline="true" :model="form" :rules="rules" label-width="120px">
                    <el-form-item label="销售单号" prop="saleNumber"> 
                        <el-select v-model="form.saleNumber" clearable 
                                         @change="getSaleListGoods"
                                         placeholder="请选择" >
                                <el-option
                                        v-for="dict in saleNumberOptions"
                                        :key="dict.key"
                                        :label="dict.value"
                                        :value="dict.key"
                                />
                            </el-select>
                    </el-form-item>
                    <el-form-item label="采购日期" prop="purchaseDate">
                        <el-date-picker v-model="form.purchaseDate" format="yyyy-MM-dd" value-format="yyyy-MM-dd"
                        :style="{width: '100%'}" placeholder="采购日期" clearable></el-date-picker>
                    </el-form-item>
                    <el-form-item label="应付金额(￥)" prop="amountPayable"> 
                        <el-input-number v-model="form.amountPayable" placeholder="应付金额" :disabled='true' :step="0.01"
                                        />
                    </el-form-item>
                    <el-form-item label="实付付金额(￥)" prop="amountPaid">
                        <el-input-number v-model="form.amountPaid" placeholder="实付金额" :step="0.01" :min="0"
                                        />
                    </el-form-item>
                    <el-form-item label="是否付款" prop="state">
                        <el-select v-model="form.state"
                                        placeholder="请选择" >
                                <el-option
                                        v-for="dict in stateOptions"
                                        :key="dict.value"
                                        :label="dict.label"
                                        :value="dict.value"
                                />
                            </el-select>
                    </el-form-item>
                    <el-form-item label="备注" prop="remarks">
                        <el-input v-model="form.remarks" placeholder="备注"
                                        />
                    </el-form-item>
                    </br>
                    <h3>相关商品列表:</h3>
                    <el-table v-loading="goodsLoading" :data="saleListGoodsList">
                    <el-table-column  width="55" align="center"/><el-table-column label="商品编码" align="center" prop="gCode"
                            :show-overflow-tooltip="true"/><el-table-column label="商品名称" align="center" prop="gName"
                            :show-overflow-tooltip="true"/><el-table-column label="商品规格" align="center" prop="specification"
                            :show-overflow-tooltip="true"/><el-table-column label="单位" align="center" prop="unit"
                            :show-overflow-tooltip="true"/><el-table-column label="数量" align="center" prop="num"
                            :show-overflow-tooltip="true"/><el-table-column label="上次采购价(￥)" align="center" prop="lastPurchasingPrice"
                            :show-overflow-tooltip="true"/>
                            
                            <el-table-column label="本次采购价(￥)" align="center" prop="price"
                            :show-overflow-tooltip="true">
                            <template slot-scope="scope">
                                <el-input
                                        v-model.number="scope.row.price"
                                        placeholder="输入采购价"
                                        type="number"
                                        step="0.01"  
                                        clearable
                                        @change="updateTotalPrice(scope.row)"
                                        :min="0">
                                        
                                    </el-input>  
                            </template>
                            </el-table-column>
                            <el-table-column label="总价(￥)" align="center" prop="totalPrice"
                            :show-overflow-tooltip="true"/><el-table-column label="商品" v-if="false" align="center" prop="goodsId"
                            :show-overflow-tooltip="true"/>
                        <!-- 新增供应商下拉框列 -->
                        <el-table-column label="供货商" align="center" width="150" prop="supplierId">
                            <template slot-scope="scope">
                                <el-select v-model="scope.row.supplierId" placeholder="选择供应商">
                                    <el-option
                                        v-for="supplier in supplierIdOptions"
                                        :key="supplier.key"
                                        :label="supplier.value"
                                        :value="supplier.key">
                                    </el-option>
                                </el-select>
                            </template>
                        </el-table-column>

                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                slot="reference"
                                size="mini"
                                type="text"
                                icon="el-icon-edit"
                            @click="removeItem()"
                            >修改
                            </el-button>
                            <el-button
                                slot="reference"
                                size="mini"
                                type="text"
                                icon="el-icon-delete"
                            @click="removeItem()"
                            >删除
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
                    <el-form-item size="small">
                        <el-button type="primary" @click="submitForm">保存</el-button>
                        <el-button @click="reset">重置</el-button>
                        </el-form-item>
                </el-form>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {addPurchaseList, getPurchaseList, listPurchaseList, updatePurchaseList} from '@/api/jxc/tb-purchase-list'
    import { getSaleList, listSaleList} from '@/api/jxc/tb-sale-list'
    import {listSaleListGoods} from '@/api/jxc/tb-sale-list-goods'
    import {listSupplier } from '@/api/jxc/tb-supplier'
    import {getGoods, listGoods} from '@/api/jxc/tb-goods'
    
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
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                purchaseListList: [],
                stateOptions: [],
                storageStatusOptions: [],
                // 关系表类型
                saleNumberOptions: [],
                saleListGoodsList: [],
                supplierIdOptions:[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10000,
                    shipmentStatus: "0",
                    createdAtOrder: "desc",
                    isPurchased: "0"
                },
                // 表单参数
                form: {
                    state:"1"
                },
                goodsLoading: false,
                goodsQueryParams:{
                    pageIndex: 1,
                    pageSize: 10000,
                    saleListId: undefined,
                },
                // 表单校验
                 rules: {
                    state:  [ {required: true, message: '是否付款不能为空', trigger: 'blur'} ],
                    purchaseDate:  [ {required: true, message: '销售日期不能为空', trigger: 'blur'} ],
                    saleNumber:  [ {required: true, message: '销售订单不能为空', trigger: 'blur'} ],

                },
                selectedSaleNumberValue: '' // 用于存储选中项的 value
        }
        },
        created() {
            this.getDicts('bus_sale_list_status').then(response => {
                this.stateOptions = response.data
            })
            this.getDicts('bus_purchase_list_storage_status').then(response => {
                this.storageStatusOptions = response.data
            })
            this.getSaleListOptionsItems()
            this.getSupplierItems()
            },
        methods: {
            // 销售单关系
            saleNumberFormat(row) {
                return this.selectItemsLabel(this.saleNumberOptions, row.saleNumber)
            },
             // 关系
             getSaleListOptionsItems() {
               this.getItems(listSaleList, this.queryParams).then(res => {
                   this.saleNumberOptions = this.setItems(res, 'id', 'saleNumber')
               })
            },
            // 关系
            getSupplierItems() {
               this.getItems(listSupplier, undefined).then(res => {
                   this.supplierIdOptions = this.setItems(res, 'id', 'supplierName')
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
                storageStatus: undefined,
                state: "1",
                customerId: undefined,
                saleNumber:undefined,
            }
                this.resetForm('form')
                this.saleListGoodsList = []
               // this.supplierIdOptions = [] // 确保供应商选项清空
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            supplierIdFormat(row) {
                return this.selectItemsLabel(this.supplierIdOptions, row.supplierId)
            },
            stateFormat(row) {
                return this.selectDictLabel(this.stateOptions, row.state)
            },
            storageStatusFormat(row) {
                return this.selectDictLabel(this.storageStatusOptions, row.storageStatus)
            },
            
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
                this.title = '添加查询采购单'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.id)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
          
            /**销售商品明细列表 */
            getSaleListGoods(id) {
                if (!id) {
                    // 当 id 为空时，清空列表数据
                    this.reset()
                    this.goodsLoading = false
                    return;
                 }
                const selectedOption = this.saleNumberOptions.find(item => item.key === id);
                if (selectedOption) {
                    this.selectedSaleNumberValue = selectedOption.value;
                    console.log('选中的销售单号值为:', this.selectedSaleNumberValue);
                 }
                
                this.goodsQueryParams.saleListId = id
                this.goodsLoading = true
                listSaleListGoods(this.goodsQueryParams).then(response => {
                        this.saleListGoodsList = response.data.list
                        this.saleListGoodsList.forEach(item => {
                            getGoods(item.goodsId).then(response => {
                                item.lastPurchasingPrice = response.data.lastPurchasingPrice
                                item.price = response.data.lastPurchasingPrice
                                this.updateTotalPrice(item)
                            })
                         });
                        this.calculateTotalPayable() // 计算应付金额    
                        this.calculateTotalPaid() // 计算实付金额
                        this.goodsLoading = false
                    })
                .catch(error => {
                    console.error('获取销售单商品列表失败:', error);
                    this.goodsLoading = false;
                });
            },
            
             // 计算临时商品数组一行总金额
            updateTotalPrice(row) {
                row.totalPrice = parseFloat(( row.price * row.num)); // 保留两位小数
                this.calculateTotalPayable() // 计算应付金额    
                this.calculateTotalPaid() // 计算实付金额
            },
            // 计算应付金额
            calculateTotalPayable() {
                const total = this.saleListGoodsList.reduce((sum, item) => sum + item.totalPrice, 0);
                this.form.amountPayable = total; // 保留两位小数
             },
             // 计算实付金额
             calculateTotalPaid() {
                if (this.form.state === "1") {
                    this.form.amountPaid = this.form.amountPayable ; // 保留两位小数
                }
            },
            // 删除临时数组单条数据
            removeItem() {
                alert("功能待开发")
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    this.saleListGoodsList.forEach(item => {
                        item.supplierId = Number(item.supplierId); // 转换为数字
                    });
                    const payload = {
                        ...this.form,
                        goods: this.saleListGoodsList // 将对象数组作为表单的一部分
                    };
                    payload.selectedSaleNumberValue =  this.selectedSaleNumberValue

                    if (valid) {
                        if (this.form.id !== undefined) {
                            updatePurchaseList(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.reset()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addPurchaseList(payload).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.reset()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        }
                    }
                })
            }
           
        }
    }
</script>
