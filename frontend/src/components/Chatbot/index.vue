<template>
  <div>
    <VueBotUI
      :messages="data"
      :options="botOptions"
      @msg-send="messageSendHandler"
    />
  </div>
</template>

<script>
import { VueBotUI } from 'vue-bot-ui'
import { wsStock } from '@/api/stock'

let ws = null

export default {
  name: 'Chatbot',
  components: {
    VueBotUI
  },
  data() {
    return {
      loading: false,
      type: '',
      data: [
        {
          agent: 'bot',
          type: 'button',
          text: '請選擇詢問項目',
          disableInput: true,
          options: [
            {
              text: 'Stock',
              value: 'stock',
              action: 'postback' // Request to API
            }
          ]
        }
      ],
      botOptions: {
      }
    }
  },
  computed: {
  },
  async created() {
    this.loading = true
    try {
      ws = await wsStock()
      ws.onmessage = async({ data }) => {
        switch (this.type) {
          case 'stock': {
            const list = JSON.parse(data)
            const msg = `
            成交量: ${list.data[0].deal_count} \n
            收盤價: ${list.data[0].price_on_close} \n
            最高價: ${list.data[0].price_on_highest} \n
            最低價: ${list.data[0].price_on_lowest} \n
            開盤價: ${list.data[0].price_on_open} \n
            `
            this.data.push({
              agent: 'bot',
              type: 'text',
              text: msg,
              disableInput: false
            })
            break
          }
        }
        this.loading = false
      }
    } catch (error) {
      this.loading = false
    }
  },
  destroyed() {
    if (ws != null) {
      ws.close()
    }
  },
  methods: {
    async messageSendHandler(msg) {
      switch (msg.action) {
        case 'postback':
          this.type = msg.value
          switch (this.type) {
            case 'stock':
              this.data.push({
                agent: 'bot',
                type: 'text',
                text: '請輸入股票編號與日期(0050,2021-01-12)',
                disableInput: false
              })
              break
          }
          break
        // 代表為User傳送的
        default:
          switch (this.type) {
            case 'stock': {
              const data = msg.text.split(',')
              if (data.length === 2) {
                ws.send(JSON.stringify({
                  type: this.type,
                  data: {
                    'stock_number': data[0],
                    'start_date': this.moment.utc(data[1]),
                    'end_date': this.moment.utc(data[1]).add(1, 'days')
                  }
                }))
              }
              break
            }
          }
          this.data.push({
            agent: 'user',
            type: 'text',
            text: msg.text,
            disableInput: false
          })
          break
      }
    }
  }
}
</script>
