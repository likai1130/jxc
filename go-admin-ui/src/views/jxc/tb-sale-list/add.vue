
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="form" :inline="true" :model="form" :rules="rules" label-width="120px">
                    <el-form-item label="客户" prop="customerId">
                        <el-select v-model="form.customerId" clearable
                                        placeholder="请选择" >
                                <el-option
                                        v-for="dict in customerIdOptions"
                                        :key="dict.key"
                                        :label="dict.value"
                                        :value="dict.key"
                                />
                            </el-select>
                    </el-form-item>
                    <el-form-item label="销售日期" prop="saleDate">
                        <el-date-picker v-model="form.saleDate" format="yyyy-MM-dd" value-format="yyyy-MM-dd"
                            :style="{width: '100%'}" placeholder="销售日期" clearable></el-date-picker>
                    </el-form-item>
                    <el-form-item label="应付金额(￥)" prop="amountPayable"> 
                        <el-input-number v-model="form.amountPayable" placeholder="应付金额" :disabled='true' :step="0.01"
                                        />
                    </el-form-item>
                    <el-form-item label="实付金额(￥)" prop="amountPaid">
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
                    <el-row :gutter="10" class="mb8">
                        <el-col :span="1.5">
                            <el-button
                                    v-permisaction="['jxc:goods:list']"
                                    type="primary"
                                    icon="el-icon-plus"
                                    size="mini"
                                    @click="handleAdd"
                            >添加商品明细
                            </el-button>
                        </el-col>
                    </el-row>
                    <!-- 临时表单 -->
                    <el-table v-loading="loading" :data="tempItem">
                        <el-table-column width="55" align="center"/><el-table-column label="商品编码" align="center" prop="gCode"
                                                    :show-overflow-tooltip="true"/><el-table-column label="商品名称" align="center" prop="gName"
                                                    :show-overflow-tooltip="true"/><el-table-column label="商品规格" align="center" prop="specification"
                                                    :show-overflow-tooltip="true"/><el-table-column label="单位" align="center" prop="unit"  width="100">
                                </el-table-column>
                                <el-table-column label="单价(￥)" align="center" prop="price"
                                                    :show-overflow-tooltip="true"/><el-table-column label="数量" align="center" prop="num"
                                                    :show-overflow-tooltip="true"/><el-table-column label="总金额(￥)" align="center" prop="totalPrice"
                                                    :show-overflow-tooltip="true"/><el-table-column label="商品ID" v-if="false" align="center" prop="goodsId"
                                                    :show-overflow-tooltip="true"/>
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
                <!-- 商品列表对话框 -->
                <el-dialog :title="title" :visible.sync="open">
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

                        <el-table v-loading="loading" :data="goodsList">
                            <el-table-column width="55" align="center"/><el-table-column label="商品编码" align="center" prop="gCode"
                                                        :show-overflow-tooltip="true"/><el-table-column label="商品名称" align="center" prop="gName"
                                                        :show-overflow-tooltip="true"/><el-table-column label="商品规格" align="center" prop="specification"
                                                        :show-overflow-tooltip="true"/><el-table-column label="单位" align="center" prop="unit">
                                    
                                    </el-table-column><el-table-column label="上次进价(￥)" align="center" prop="lastPurchasingPrice"
                                                        :show-overflow-tooltip="true"/><el-table-column label="销售价(￥)" align="center" prop="sellingPrice"
                                                        :show-overflow-tooltip="true"/><el-table-column label="安全库存量" align="center" prop="safetyStock"
                                                        :show-overflow-tooltip="true"/><el-table-column label="商品描述" align="center" prop="description"
                                                        :show-overflow-tooltip="true"/><el-table-column label="生产厂商" align="center" prop="producer"
                                                        :show-overflow-tooltip="true"/>
                            <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                                <template slot-scope="scope">
                                    <el-button
                                        slot="reference"
                                        size="mini"
                                        type="text"
                                        icon="el-icon-goods"
                                        @click="handleUpdate(scope.row)"
                                    >选择
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

                        <!-- 商品详情对话框 -->
                        <el-dialog :title="goodsTitle" :visible.sync="goodsOpen" :modal-append-to-body="false" :append-to-body="true">
                            <el-form ref="formGoods" :label-position="labelPosition" :model="formGoods" :rules="rules" label-width="100px">
                                
                                <el-form-item label="商品编码" prop="gCode">
                                    <el-input v-model="formGoods.gCode" placeholder="商品编码自动生成" :disabled="true"
                                                    />
                                </el-form-item>
                                <el-form-item label="商品名称" prop="gName">
                                    <el-input v-model="formGoods.gName" placeholder="商品名称" :disabled='true'
                                                    />
                                </el-form-item>
                                <el-form-item label="商品规格" prop="specification">
                                    <el-input v-model="formGoods.specification" placeholder="商品规格" :disabled='true'
                                                    />
                                </el-form-item>
                                <el-form-item label="单位" prop="unit">
                                    <el-select v-model="formGoods.unit" :disabled='true'
                                                    placeholder="请选择" >
                                            <el-option
                                                    v-for="dict in unitOptions"
                                                    :key="dict.key"
                                                    :label="dict.value"
                                                    :value="dict.key"
                                            />
                                        </el-select>
                                </el-form-item>
                                <el-form-item label="上次进价(￥)" prop="lastPurchasingPrice">
                                    <el-input-number v-model="formGoods.lastPurchasingPrice" placeholder="上次进价价" :disabled='true' :step="0.01"
                                                    />
                                </el-form-item>
                                <el-form-item label="当前库存量" prop="stockQuantity">
                                    <el-input-number v-model="formGoods.stockQuantity" placeholder="当前库存量" :disabled='true'
                                                    />
                                </el-form-item>
                                <el-form-item label="安全库存量" prop="safetyStock">
                                    <el-input-number v-model="formGoods.safetyStock" placeholder="安全库存量" :disabled='true'
                                                    />
                                </el-form-item>
                                <el-form-item label="商品描述" prop="description">
                                    <el-input v-model="formGoods.description" placeholder="商品描述" :disabled='true'
                                                    />
                                </el-form-item>
                                <el-form-item label="生产厂商" prop="producer">
                                    <el-input v-model="formGoods.producer" placeholder="生产厂商" :disabled='true'
                                                    />
                                </el-form-item>

                                <el-form-item label="销售价(￥)" prop="sellingPrice">
                                    <el-input-number v-model="formGoods.sellingPrice" :step="0.01" placeholder="销售指导价"/>
                                </el-form-item>

                                <el-form-item label="数量" prop="num">
                                <el-input-number v-model="formGoods.num" :min="1" placeholder="请输入数量"></el-input-number>
                                </el-form-item>
                            </el-form>
                            <div slot="footer" class="dialog-footer">
                                <el-button type="primary" @click="addItem">确 定</el-button>
                                <el-button @click="cancelGoods">取 消</el-button>
                            </div>
                     </el-dialog>
                    </el-card>
                </el-dialog>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {addSaleList, delSaleList, getSaleList, listSaleList, updateSaleList} from '@/api/jxc/tb-sale-list'
    import {getGoods, listGoods} from '@/api/jxc/tb-goods'
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
                open: false,
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                saleListList: [],
                stateOptions: [],
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
                    state:"1"
                },
                // 表单校验
                rules: {
                    state:  [ {required: true, message: '状态不能为空', trigger: 'blur'} ],
                    customerId:  [ {required: true, message: '客户不能为空', trigger: 'change'} ],
                    saleDate:  [ {required: true, message: '销售日期不能为空', trigger: 'blur'} ],
                },
                //--------以下是嵌套表单参数
                labelPosition: 'left',
                formGoods:{},
                 // 嵌套的商品列表
                goodsList: [],
                  // 关系表类型
                unitOptions :[],
                goodsOpen: false,
                goodsTitle: '',
                tempItem: [],

        }
        },
        created() {
            this.getList()
            this.getDicts('bus_sale_list_status').then(response => {
                this.stateOptions = response.data
            })
            this.getConsumerItems()
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listGoods(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.goodsList = response.data.list
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
                state: "1",
                customerId: undefined,
            }
                this.resetForm('form')
                this.tempItem = []
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
            /** 新增按钮操作 */
            handleAdd() {
               // this.reset()
                this.open = true
                this.title = '添加商品明细'
                this.isEdit = false
            },
       
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                     this.form.customerId = parseInt(this.form.customerId, 10);
                     const payload = {
                        ...this.form,
                        goods: this.tempItem // 将对象数组作为表单的一部分
                    };

                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateSaleList(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.reset()
                                    //this.open = false
                                    //this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {

                            addSaleList(payload).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.reset()
                                    //this.open = false
                                    //this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        }
                    }
                })
            },
            /**---------------------------------------------------------- */
            /** 商品列表  表单重置 */
             resetGoods() {
                this.formGoods = {
                
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
                this.resetForm('formGoods')
            },

            /** 商品列表  单位 */
            unitFormat(row) {
                return this.selectItemsLabel(this.unitOptions, row.unit)
            },
             // 关系
            getGoodsUnitItems() {
               this.getItems(listGoodsUnit, undefined).then(res => {
                   this.unitOptions = this.setItems(res, 'name', 'name')
               })
            },
            // 选择商品操作
            handleUpdate(row) {
                this.resetGoods()
                const id =
                row.id || this.ids
                getGoods(id).then(response => {
                    this.formGoods = response.data
                    this.goodsOpen = true
                    this.goodsTitle = '商品详情'
                    this.isEdit = true
                })
            },
             // 取消按钮
             cancelGoods() {
                this.goodsOpen = false
                this.resetGoods()
            },
            // 添加商品到数据组
            addItem() {
                // 将表单数据复制到临时数组中
                this.updateTotalPrice();
                this.tempItem.push({ ...this.formGoods });
                this.cancelGoods(); // 关闭对话框
                this.calculateTotalPayable(); //计算应付金额
                this.calculateTotalPaid(); // 计算实付金额
                
            },
            // 计算临时商品数组一行总金额
            updateTotalPrice() {
                this.formGoods.goodsId = this.formGoods.id
                this.formGoods.price = this.formGoods.sellingPrice
                const unitPrice = parseFloat(this.formGoods.sellingPrice) || 0; // 确保是浮点数
                const quantity = this.formGoods.num || 0;   // 确保数量为数字
                this.formGoods.totalPrice = parseFloat((unitPrice * quantity)); // 保留两位小数
            },
            // 计算应付金额
            calculateTotalPayable() {
                const total = this.tempItem.reduce((sum, item) => sum + item.totalPrice, 0);
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
        }
    }
</script>
