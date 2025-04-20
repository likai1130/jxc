
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="销售单号" prop="saleNumber"><el-input v-model="queryParams.saleNumber" placeholder="请输入销售单号" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="是否付款" prop="state"><el-select v-model="queryParams.state"
                                               placeholder="是否付款" clearable size="small">
                                        <el-option
                                                v-for="dict in stateOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                        </el-form-item>
                        <el-form-item label="是否采购" prop="isPurchased"><el-select v-model="queryParams.isPurchased"
                                               placeholder="是否采购" clearable size="small">
                                        <el-option
                                                v-for="dict in isPurchasedOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                        </el-form-item>
                        <el-form-item label="是否出库" prop="shipmentStatus"><el-select v-model="queryParams.shipmentStatus"
                                               placeholder="是否出库" clearable size="small">
                                        <el-option
                                                v-for="dict in shipmentStatusOptions"
                                                :key="dict.value"
                                                :label="dict.label"
                                                :value="dict.value"
                                        />
                                    </el-select>
                            </el-form-item>
    
                        <el-form-item label="客户" prop="customerId"><el-select v-model="queryParams.customerId"
                                           placeholder="请选择" clearable size="small" >
                                    <el-option
                                            v-for="dict in customerIdOptions"
                                            :key="dict.key"
                                            :label="dict.value"
                                            :value=dict.key
                                    />
                                </el-select>
                            </el-form-item>
                        
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-table v-loading="loading" :data="saleListList">
                    <el-table-column width="55" align="center"/><el-table-column label="销售单号" align="center" prop="saleNumber"
                    :show-overflow-tooltip="true"/><el-table-column label="总金额(￥)" align="center" prop="amountPayable"
                    :show-overflow-tooltip="true"/><el-table-column label="已付金额(￥)" align="center" prop="amountPaid"
                    :show-overflow-tooltip="true"/><el-table-column label="销售日期" align="center" prop="saleDate"
                                                 :show-overflow-tooltip="true"/>
                        <el-table-column label="客户" align="center" prop="customerId" :formatter="customerIdFormat" width="150">
                                    <template slot-scope="scope">
                                        {{ customerIdFormat(scope.row) }}
                                    </template>
                                </el-table-column>
                        <el-table-column label="备注" align="center" prop="remarks"
                                        :show-overflow-tooltip="true"/><el-table-column label="是否付款" align="center" prop="state"
                                        :formatter="stateFormat" width="100">
                        <template slot-scope="scope">
                            {{ stateFormat(scope.row) }}
                        </template>
                    </el-table-column>
                    <el-table-column label="是否采购" align="center" prop="isPurchased"
                                        :formatter="isPurchasedFormat" width="100">
                        <template slot-scope="scope">
                            {{ isPurchasedFormat(scope.row) }}
                        </template>
                    </el-table-column>
                    <el-table-column label="是否出库" align="center" prop="shipmentStatus"
                                        :formatter="shipmentStatusFormat" width="100">
                        <template slot-scope="scope">
                            {{ shipmentStatusFormat(scope.row) }}
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
                             size="mini"
                             type="text"
                             icon="el-icon-s-goods"
                            @click="handleOpenTable(scope.row)"
                           >销售明细
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
                    title="销售单商品信息"
                    :visible.sync="openTable"
                    direction="rtl"
                    size="50%">
                    <p>销售单号: <span v-text="currentRow.saleNumber"></span></p>
                    <br>
                    <el-table v-loading="goodsLoading" :data="saleListGoodsList">
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
    import {addSaleList, delSaleList, getSaleList, listSaleList, updateSaleList} from '@/api/jxc/tb-sale-list'
    import {listSaleListGoods} from '@/api/jxc/tb-sale-list-goods'
    
    import {listConsumer } from '@/api/jxc/tb-consumer'
    export default {
        name: 'SaleList',
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
                openTable: false,
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                saleListList: [],
                stateOptions: [],
                shipmentStatusOptions:[],
                isPurchasedOptions:[],
                // 关系表类型
                customerIdOptions :[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    saleNumber:undefined,
                    state:undefined,
                    customerId:undefined,
                    createdAtOrder: "desc"
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {},
                // 商品明细列表
                saleListGoodsList: [],
                goodsLoading: false,
                goodsQueryParams:{
                    pageIndex: 1,
                    pageSize: 10,
                    saleListId: undefined,
                },
                currentRow: {}, // 当前选中的行数据
                goodsTotal: 0 // 商品总条数
        }
        },
        created() {
            this.getList()
            this.getDicts('bus_sale_list_status').then(response => {
                this.stateOptions = response.data
            })
            this.getDicts('bus_sale_list_shipment_status').then(response => {
                this.shipmentStatusOptions = response.data
            })
            this.getDicts('bus_sale_list_is_purchased').then(response => {
                this.isPurchasedOptions = response.data
            })
            this.getConsumerItems()
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listSaleList(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.saleListList = response.data.list
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
                saleDate: undefined,
                remarks: undefined,
                state: undefined,
                customerId: undefined,
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
            
            shipmentStatusFormat(row) {
                return this.selectDictLabel(this.shipmentStatusOptions, row.shipmentStatus)
            },

            isPurchasedFormat(row) {
                return this.selectDictLabel(this.isPurchasedOptions, row.isPurchased)
            },

            customerIdFormat(row) {
                return this.selectItemsLabel(this.customerIdOptions, row.customerId)
            },
            // 关系
            getConsumerItems() {
               this.getItems(listConsumer, undefined).then(res => {
                   this.customerIdOptions = this.setItems(res, 'id', 'customerName')
               })
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
            /** 打开商品明细操作 */
            handleOpenTable(row) {
                this.currentRow = { ...row }; // 更新当前行数据
                this.openTable = true
                this.getSaleListGoods(row.id)

            },
            /**商品明细列表 */
            getSaleListGoods(id) {
                this.goodsQueryParams.saleListId = id
                this.goodsLoading = true
                listSaleListGoods(this.goodsQueryParams).then(response => {
                        this.saleListGoodsList = response.data.list
                        this.goodsTotal = response.data.count
                        this.goodsLoading = false
                    }
                )
            }
        }
    }
</script>
