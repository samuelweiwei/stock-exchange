
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="the ID of the stock:" prop="wid">
          <el-input v-model.number="formData.wid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="the name of the stock:" prop="widCode">
          <el-input v-model="formData.widCode" :clearable="true"  placeholder="请输入the name of the stock" />
       </el-form-item>
        <el-form-item label="0 Flexible 活期, 1 Fixed 定期:" prop="type">
          <el-switch v-model="formData.type" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="产品名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入产品名称" />
       </el-form-item>
        <el-form-item label="当前利率:" prop="currentInterestRates">
          <el-input-number v-model="formData.currentInterestRates" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="最小利率:" prop="minInterestRates">
          <el-input-number v-model="formData.minInterestRates" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="最大利率:" prop="maxInterestRates">
          <el-input-number v-model="formData.maxInterestRates" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="模拟|||默认不是模拟:" prop="isMoni">
          <el-switch v-model="formData.isMoni" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="违约金比例:" prop="penaltyRatio">
          <el-input-number v-model="formData.penaltyRatio" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="product marks:" prop="mark">
          <el-input v-model="formData.mark" :clearable="true"  placeholder="请输入product marks" />
       </el-form-item>
        <el-form-item label="the stock of the product -1 无限库存:" prop="stock">
          <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="the days of the projects 理财定期时长:" prop="duration">
          <el-input v-model.number="formData.duration" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态|||1:可用,2:冻结:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="createdAt字段:" prop="createdAt">
          <el-input v-model.number="formData.createdAt" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="updatedAt字段:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="后台管理员备注:" prop="adminRemark">
          <el-input v-model="formData.adminRemark" :clearable="true"  placeholder="请输入后台管理员备注" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createEarnProducts,
  updateEarnProducts,
  findEarnProducts
} from '@/api/earn/earnProducts'

defineOptions({
    name: 'EarnProductsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: undefined,
            wid: undefined,
            widCode: '',
            type: false,
            name: '',
            currentInterestRates: 0,
            minInterestRates: 0,
            maxInterestRates: 0,
            isMoni: false,
            penaltyRatio: 0,
            mark: '',
            stock: undefined,
            duration: undefined,
            status: false,
            createdAt: undefined,
            updatedAt: new Date(),
            adminRemark: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEarnProducts({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createEarnProducts(formData.value)
               break
             case 'update':
               res = await updateEarnProducts(formData.value)
               break
             default:
               res = await createEarnProducts(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
