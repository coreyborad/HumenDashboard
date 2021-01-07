<template>
  <el-autocomplete
    v-model="val"
    style="width: 100%;"
    class="inline-input"
    :fetch-suggestions="querySearch"
    :disabled="disabled"
    placeholder="請輸入內容"
    @input="valueChange"
  />
</template>

<script>
import { getMakeupByQuery } from '@/api/makeup'

export default {
  props: {
    value: {
      type: String,
      required: true,
      default: ''
    },
    brand: {
      type: String,
      required: false,
      default: ''
    },
    target: {
      type: String,
      required: true,
      default: 'brand'
    },
    disabled: {
      type: Boolean,
      required: true,
      default: false
    }
  },
  data() {
    return {
      val: this.value
    }
  },
  computed: {
  },
  async created() {
  },
  methods: {
    // todo 更換form
    async querySearch(queryString, cb) {
      if (queryString.length >= 2) {
        let qs = {}
        let data = []
        switch (this.target) {
          case 'brand': {
            qs = {
              brand: queryString
            }
            const temp = await getMakeupByQuery(qs)
            data = temp.map(r => { return { value: r.brand } }).filter((v, i, a) => a.findIndex(t => (t.value === v.value)) === i)
            break
          }
          case 'name': {
            qs = {
              brand: this.brand,
              name: queryString
            }
            const temp = await getMakeupByQuery(qs)
            data = temp.map(r => { return { value: r.name } })
            break
          }
        }
        cb(data)
      } else {
        cb([])
      }
    },
    valueChange(val) {
      this.$emit('change', val)
    }
  }
}
</script>
