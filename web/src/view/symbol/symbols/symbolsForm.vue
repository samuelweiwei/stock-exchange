
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主键:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="股票代码，通常是股票的唯一标识:" prop="symbol">
          <el-input v-model="formData.symbol" :clearable="true"  placeholder="请输入股票代码，通常是股票的唯一标识" />
       </el-form-item>
        <el-form-item label="公司名称:" prop="corporation">
          <el-input v-model="formData.corporation" :clearable="true"  placeholder="请输入公司名称" />
       </el-form-item>
        <el-form-item label="行业:" prop="industry">
          <el-input v-model="formData.industry" :clearable="true"  placeholder="请输入行业" />
       </el-form-item>
        <el-form-item label="所属交易所:" prop="exchange">
          <el-input v-model="formData.exchange" :clearable="true"  placeholder="请输入所属交易所" />
       </el-form-item>
        <el-form-item label="市值:" prop="marketCap">
          <el-input v-model.number="formData.marketCap" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="IPO 上市日期:" prop="ipoDate">
          <el-date-picker v-model="formData.ipoDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="最新价格:" prop="currentPrice">
          <el-input-number v-model="formData.currentPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="日均交易量:" prop="averageVolume">
          <el-input v-model.number="formData.averageVolume" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="日涨跌幅度:" prop="changeRatio">
          <el-input-number v-model="formData.changeRatio" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="市盈率:" prop="peRatio">
          <el-input-number v-model="formData.peRatio" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="更新时间:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createSymbols,
  updateSymbols,
  findSymbols
} from '@/api/symbol/symbols'

defineOptions({
    name: 'SymbolsForm'
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
            symbol: '',
            corporation: '',
            industry: '',
            exchange: '',
            marketCap: undefined,
            ipoDate: new Date(),
            currentPrice: 0,
            averageVolume: undefined,
            changeRatio: 0,
            peRatio: 0,
            createdAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSymbols({ ID: route.query.id })
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
               res = await createSymbols(formData.value)
               break
             case 'update':
               res = await updateSymbols(formData.value)
               break
             default:
               res = await createSymbols(formData.value)
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
