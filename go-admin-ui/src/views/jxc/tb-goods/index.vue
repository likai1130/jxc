
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

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['jxc:goods:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['jxc:goods:edit']"
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
                                v-permisaction="['jxc:goods:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="goodsList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="商品编码" align="center" prop="gCode"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品名称" align="center" prop="gName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品规格" align="center" prop="specification"
                                                 :show-overflow-tooltip="true"/><el-table-column label="单位" align="center" prop="unit" :formatter="unitFormat" width="100">
                                <template slot-scope="scope">
                                    {{ unitFormat(scope.row) }}
                                </template>
                            </el-table-column><el-table-column label="采购价(￥)" align="center" prop="purchasingPrice"
                                                 :show-overflow-tooltip="true"/><el-table-column label="销售价(￥)" align="center" prop="sellingPrice"
                                                 :show-overflow-tooltip="true"/><el-table-column label="安全库存量" align="center" prop="safetyStock"
                                                 :show-overflow-tooltip="true"/><el-table-column label="商品描述" align="center" prop="description"
                                                 :show-overflow-tooltip="true"/><el-table-column label="生产厂商" align="center" prop="producer"
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
                             v-permisaction="['jxc:goods:edit']"
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
                              v-permisaction="['jxc:goods:remove']"
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
                    <el-form ref="form" :label-position="labelPosition" :model="form" :rules="rules" label-width="100px">
                        
                                    <el-form-item label="商品编码" prop="gCode">
                                        <el-input v-model="form.gCode" placeholder="商品编码自动生成" :disabled="true"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品名称" prop="gName">
                                        <el-input v-model="form.gName" placeholder="商品名称"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品规格" prop="specification">
                                        <el-input v-model="form.specification" placeholder="商品规格"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="单位" prop="unit">
                                        <el-select v-model="form.unit"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in unitOptions"
                                                        :key="dict.key"
                                                        :label="dict.value"
                                                        :value="dict.key"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="采购价(￥)" prop="purchasingPrice">
                                        <el-input-number v-model="form.purchasingPrice" placeholder="采购价" :step="0.01" :min="0"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="销售价(￥)" prop="sellingPrice">
                                        <el-input-number v-model="form.sellingPrice" placeholder="销售指导价"  :step="0.01" :min="0"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="安全库存量" prop="safetyStock">
                                        <el-input-number v-model="form.safetyStock" placeholder="安全库存量" :min="0"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="商品描述" prop="description">
                                        <el-input v-model="form.description" placeholder="商品描述"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="生产厂商" prop="producer">
                                        <el-input v-model="form.producer" placeholder="生产厂商"
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
    import {addGoods, delGoods, getGoods, listGoods, updateGoods} from '@/api/jxc/tb-goods'
    
    import {listGoodsUnit } from '@/api/jxc/tb-goods-unit'
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
                    createdAtOrder: "desc"

                    
                },
                labelPosition: 'left',
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {
                   // gCode:  [ {required: true, message: '商品编码不能为空', trigger: 'blur'} ],
                    gName:  [ {required: true, message: '商品名称不能为空', trigger: 'blur'} ],
                    producer:  [ {required: true, message: '生产厂商不能为空', trigger: 'blur'} ],
                    specification:  [ {required: true, message: '规格不能为空', trigger: 'blur'} ],
                    unit:  [ {required: true, message: '单位不能为空', trigger: 'blur'} ],
                    purchasingPrice:  [ {required: true, message: '采购价不能为空', trigger: 'blur'} ],
                }
        }
        },
        created() {
            this.getList()
            this.getGoodsUnitItems()
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
            // 关系
            getGoodsUnitItems() {
               this.getItems(listGoodsUnit, undefined).then(res => {
                   this.unitOptions = this.setItems(res, 'name', 'name')
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
                this.reset()
                this.open = true
                this.title = '添加商品管理'
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
                getGoods(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改商品管理'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateGoods(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addGoods(this.form).then(response => {
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
                      return delGoods( { 'ids': Ids })
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
