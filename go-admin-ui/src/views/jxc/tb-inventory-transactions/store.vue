
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="商品编码" prop="gCode"><el-input v-model="queryParams.gCode" placeholder="请输入商品编码"  clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="商品名称" prop="gName"><el-input v-model="queryParams.gName" placeholder="请输入商品名称" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="生产厂商" prop="producer"><el-input v-model="queryParams.producer" placeholder="请输入生产厂商" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>
                <el-table v-loading="loading" :data="goodsList" >
                    <el-table-column width="55" align="center"/><el-table-column label="商品编码" align="center" prop="gCode"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品名称" align="center" prop="gName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品规格" align="center" prop="specification"
                                                 :show-overflow-tooltip="true"/><el-table-column label="单位" align="center" prop="unit"
                                                 :show-overflow-tooltip="true"/><el-table-column label="生产厂商" align="center" prop="producer"
                                                 :show-overflow-tooltip="true"/><el-table-column label="库存量" align="center" prop="stockQuantity"
                                                 :show-overflow-tooltip="true"/><el-table-column label="销售总数" align="center" prop="saleNum"
                                                 :show-overflow-tooltip="true"/><el-table-column label="上次进价(￥)" align="center" prop="lastPurchasingPrice"
                                                 :show-overflow-tooltip="true"/><el-table-column label="成本均价(￥)" align="center" prop="purchasingPrice"
                                                 :show-overflow-tooltip="true"/><el-table-column label="销售价(￥)" align="center" prop="sellingPrice"
                                                 :show-overflow-tooltip="true"/><el-table-column label="库存总余额(￥)" align="center" prop="totalBalance"
                                                 :show-overflow-tooltip="true"/>
                </el-table>

                <pagination
                        v-show="total>0"
                        :total="total"
                        :page.sync="queryParams.pageIndex"
                        :limit.sync="queryParams.pageSize"
                        @pagination="getList"
                />
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {listGoods} from '@/api/jxc/tb-goods'
    
    export default {
        name: 'Goods',
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
                goodsList: [],
                
                // 关系表类型
                unitOptions :[],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    gCode:undefined,
                    gName:undefined,
                    producer:undefined,
                    
                },
                labelPosition: 'left',
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {}
        }
        },
        created() {
            this.getList()
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listGoods(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.goodsList = response.data.list
                        this.goodsList.forEach(item => {
                            item.totalBalance = (item.stockQuantity * item.purchasingPrice).toFixed(2);
                        });
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
                gCode: undefined,
                gName: undefined,
                specification: undefined,
                unit: undefined,
                purchasingPrice: undefined,
                sellingPrice: undefined,
                stockQuantity: undefined,
                safetyStock: undefined,
                description: undefined,
                producer: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            unitFormat(row) {
                return this.selectItemsLabel(this.unitOptions, row.unit)
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
            }
        }
    }
</script>
