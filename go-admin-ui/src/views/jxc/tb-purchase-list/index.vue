
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="采购单号" prop="purchaseNumber"><el-input v-model="queryParams.purchaseNumber" placeholder="请输入采购单号" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="是否付款" prop="state"><el-select v-model="queryParams.state"
                                               placeholder="查询采购单是否付款" clearable size="small">
                                        <el-option
                                                v-for="dict in stateOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="是否入库" prop="storageStatus"><el-select v-model="queryParams.storageStatus"
                                               placeholder="查询采购单是否入库" clearable size="small">
                                        <el-option
                                                v-for="dict in storageStatusOptions"
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
                <el-table v-loading="loading" :data="purchaseListList" @selection-change="handleSelectionChange">
                    <el-table-column width="55" align="center"/><el-table-column label="采购单号" align="center" prop="purchaseNumber"
                    :show-overflow-tooltip="true"/><el-table-column label="总金额(￥)" align="center" prop="amountPayable"
                    :show-overflow-tooltip="true"/><el-table-column label="已付金额(￥)" align="center" prop="amountPaid"
                    :show-overflow-tooltip="true"/><el-table-column label="采购日期" align="center" prop="purchaseDate"
                    :show-overflow-tooltip="true"/><el-table-column label="销售单号" align="center" prop="saleNumber"
                    :show-overflow-tooltip="true"/><el-table-column label="备注" align="center" prop="remarks"
                    :show-overflow-tooltip="true"/>
                    <el-table-column label="是否付款" align="center" prop="state" :formatter="stateFormat" width="100"> 
                        <template slot-scope="scope">{{ stateFormat(scope.row) }}</template>
                    </el-table-column>
                    <el-table-column label="是否入库" align="center" prop="storageStatus"
                            :formatter="storageStatusFormat" width="100">
                            <template slot-scope="scope">
                                {{ storageStatusFormat(scope.row) }}
                            </template>
                    </el-table-column>
                 
                    <el-table-column label="创建时间" align="center" prop="createdAt" width="180">
                        <template slot-scope="scope">
                            <span>{{ parseTime(scope.row.createdAt) }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                            slot="reference"
                            size="mini"
                            type="text"
                            icon="el-icon-s-goods"
                            @click="handleOpenTable(scope.row)"
                            >采购明细
                            </el-button>
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
                <!--嵌套表格-->
                <el-drawer
                    title="采购单商品信息"
                    :visible.sync="openTable"
                    direction="rtl"
                    size="50%">
                    <p>采购单号: <span v-text="currentRow.purchaseNumber"></span></p>
                    <br>
                    <el-table v-loading="goodsLoading" :data="purchaseListGoodsList">
                        <el-table-column  width="55" align="center"/><el-table-column label="商品编码" align="center" prop="gCode"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品名称" align="center" prop="gName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品规格" align="center" prop="specification"
                                                 :show-overflow-tooltip="true"/><el-table-column label="单位" align="center" prop="unit"
                                                 :show-overflow-tooltip="true"/><el-table-column label="数量" align="center" prop="num"
                                                 :show-overflow-tooltip="true"/><el-table-column label="单价(￥)" align="center" prop="price"
                                                 :show-overflow-tooltip="true"/><el-table-column label="总价(￥)" align="center" prop="totalPrice"
                                                 :show-overflow-tooltip="true"/>
                    </el-table>
                    <pagination
                            v-show="goodsTotal>0"
                            :total="goodsTotal"
                            :page.sync="goodsQueryParams.pageIndex"
                            :limit.sync="goodsQueryParams.pageSize"
                            @pagination="getSaleListGoods(currentRow.id)"
                    />
                </el-drawer>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {listPurchaseList} from '@/api/jxc/tb-purchase-list'
    import {listPurchaseListGoods} from '@/api/jxc/tb-purchase-list-goods'

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
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    purchaseNumber:undefined,
                    state:undefined,
                    storageStatus:undefined,
                    createdAtOrder: "desc"
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {purchaseNumber:  [ {required: true, message: '采购单号不能为空', trigger: 'blur'} ],
                state:  [ {required: true, message: '是否付款不能为空', trigger: 'blur'} ],
                storageStatus:  [ {required: true, message: '是否入库不能为空', trigger: 'blur'} ],
                },

                // 商品明细列表
                purchaseListGoodsList: [],
                goodsLoading: false,
                goodsQueryParams:{
                    pageIndex: 1,
                    pageSize: 10,
                    purchaseListId: undefined,
                },
                currentRow: {}, // 当前选中的行数据
                goodsTotal: 0,// 商品总条数
                // 是否显示弹出层
                openTable: false,
        }
        },
        created() {
            this.getList()
            this.getDicts('bus_sale_list_status').then(response => {
                this.stateOptions = response.data
            })
            this.getDicts('bus_purchase_list_storage_status').then(response => {
                this.storageStatusOptions = response.data
            })
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listPurchaseList(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.purchaseListList = response.data.list
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
                amountPayable: undefined,
                amountPaid: undefined,
                purchaseDate: undefined,
                remarks: undefined,
                state: undefined,
                storageStatus: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            stateFormat(row) {
                return this.selectDictLabel(this.stateOptions, row.state)
            },
            storageStatusFormat(row) {
                return this.selectDictLabel(this.storageStatusOptions, row.storageStatus)
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
                this.title = '添加查询采购单'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.id)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 打开商品明细操作 */
            handleOpenTable(row) {
                    this.currentRow = { ...row }; // 更新当前行数据
                    this.openTable = true
                this.getPurchaseListGoods(row.id)

            },
            /**商品明细列表 */
            getPurchaseListGoods(id) {
                this.goodsQueryParams.purchaseListId = id
                this.goodsLoading = true
                listPurchaseListGoods(this.goodsQueryParams).then(response => {
                        this.purchaseListGoodsList = response.data.list
                        this.goodsTotal = response.data.count
                        this.goodsLoading = false
                    }
                )
            }
        }
    }
</script>
