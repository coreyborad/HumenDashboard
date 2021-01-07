<template>
  <el-dialog :title="title" width="50%" :visible="visible" :close-on-click-modal="false" @close="close">
    <el-form ref="form" :model="form" label-position="top" :rules="rules" label-width="100px">
      <el-form-item label="品牌" prop="brand">
        <querySearch
          v-model="form.brand"
          :target="'brand'"
          :disabled="active !== 1"
          @change="val => { form.brand = val}"
        />
      </el-form-item>
      <el-form-item label="品項" prop="name">
        <querySearch
          v-model="form.name"
          :target="'name'"
          :brand="form.brand"
          :disabled="active !== 2"
          @change="val => { form.name = val}"
        />
      </el-form-item>
      <el-form-item label="色號" prop="color_name">
        <el-input v-model="form.color_name" placeholder="請輸入內容" />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button v-if="active === 3" type="primary" @click="save()">建立</el-button>
      <el-button v-if="active > 1" type="info" @click="active--">上一步</el-button>
      <el-button v-if="active < 3" type="primary" @click="next()">下一步</el-button>
      <el-button @click="close">取消</el-button>
    </div>
  </el-dialog>
</template>

<script>
import querySearch from './querySearch.vue'

export default {
  components: {
    querySearch
  },
  props: {
    visible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      loading: false,
      title: '新增品項',
      active: 1,
      rules: {
        brand: [
          { required: true, message: '請輸入文字', trigger: 'blur' },
          { min: 2, message: '請至少2個字', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '請輸入文字', trigger: 'blur' },
          { min: 2, message: '請至少2個字', trigger: 'blur' }
        ],
        color_name: [
          { required: true, message: '請輸入文字', trigger: 'blur' },
          { min: 2, message: '請至少2個字', trigger: 'blur' }
        ]
      },
      form: {
        id: 0,
        name: '',
        brand: '',
        color_name: ''
      }
    }
  },
  computed: {
  },
  async created() {
  },
  methods: {
    async next() {
      switch (this.active) {
        case 1: {
          this.$refs['form'].validateField('brand', (error) => {
            if (!error) {
              this.active++
            }
          })
          break
        }
        case 2: {
          this.$refs['form'].validateField('name', (error) => {
            if (!error) {
              this.active++
            }
          })
          break
        }
      }
    },
    async save() {
      const valid = await this.$refs['form'].validate()
      if (valid) {
        try {
          await this.$store.dispatch('makeup/createMakeupInfo', this.form)
          this.$message.success('新增成功')
          this.$emit('update:visible', false)
        } catch (error) {
          this.$message.error(error)
        }
      }
    },
    close() {
      this.$emit('update:visible', false)
    }
  }
}
</script>
