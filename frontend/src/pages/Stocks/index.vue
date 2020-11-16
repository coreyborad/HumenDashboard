<template>
  <v-data-table
    :headers="headers"
    :items="tableData"
    :items-per-page="10"
    class="elevation-1"
  >
    <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>My Stock</v-toolbar-title>
        <v-divider
          class="mx-4"
          inset
          vertical
        />
        <v-spacer />
        <v-dialog
          v-model="dialog"
          max-width="500px"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              color="primary"
              dark
              class="mb-2"
              v-bind="attrs"
              v-on="on"
            >
              New Item
            </v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="headline">Stock</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-row>
                  <v-col
                    cols="12"
                    sm="6"
                    md="4"
                  >
                    <v-text-field
                      v-model="editItem.stock_number"
                      label="Stock Number"
                    />
                  </v-col>
                  <v-col
                    cols="12"
                    sm="6"
                    md="4"
                  >
                    <v-text-field
                      v-model="editItem.shares"
                      label="Shares"
                    />
                  </v-col>
                  <v-col
                    cols="12"
                    sm="6"
                    md="4"
                  >
                    <v-text-field
                      v-model="editItem.cost"
                      label="Cost"
                    />
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer />
              <v-btn
                color="blue darken-1"
                text
                @click="close"
              >
                Cancel
              </v-btn>
              <v-btn
                color="blue darken-1"
                text
                @click="save"
              >
                Save
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-toolbar>
    </template>
    <template v-slot:[`item.income`]="{ item }">
      {{ (item.last_stock.price_on_close * item.shares) - item.cost }}
    </template>
    <template v-slot:[`item.operations`]="{ item }">
      <!-- <v-icon
        small
        class="mr-2"
        @click="editStock(item)"
      >
        mdi-pencil
      </v-icon> -->
      <v-icon
        small
        @click="deleteStock(item)"
      >
        mdi-delete
      </v-icon>
    </template>
  </v-data-table>
</template>

<script>
import {
  getUserStock, createUserStock, deleteUserStock
} from '@/api/stock'

export default {
  name: 'Stock',
  props: {
  },
  data() {
    return {
      dialog: false,
      headers: [
        {
          text: '股票編號',
          align: 'start',
          sortable: false,
          value: 'stock_number'
        },
        { text: '股票名稱', value: 'stock_info.stock_name' },
        { text: '持有股數', value: 'shares' },
        { text: '持有總成本', value: 'cost' },
        { text: '最近一次價位', value: 'last_stock.price_on_close' },
        { text: '損益', value: 'income' },
        { text: '操作', value: 'operations' }
      ],
      tableData: [],
      editItem: {
        stock_number: 0,
        cost: 0,
        shares: 0
      },
      defaultItem: {
        stock_number: 0,
        cost: 0,
        shares: 0
      }
    }
  },
  computed: {
  },
  watch: {
  },
  async created() {
    this.init()
  },
  methods: {
    async init() {
      const data = await getUserStock()
      this.tableData = data
    },
    editStock(item) {
      console.log(item)
    },
    async deleteStock(item) {
      await deleteUserStock(item.id)
      this.init()
    },
    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editItem = Object.assign({}, this.defaultItem)
      })
    },
    async save() {
      await createUserStock(this.editItem)
      this.dialog = false
      this.init()
    }
  }
}
</script>
