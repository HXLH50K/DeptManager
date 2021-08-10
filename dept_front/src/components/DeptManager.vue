<template>
  <div>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column fixed prop="ID" label="部门编号" width="150">
      </el-table-column>
      <el-table-column prop="Dname" label="部门名称" width="120">
      </el-table-column>
      <el-table-column prop="Loc" label="地区" width="120"> </el-table-column>
      <el-table-column width="100">
        <template slot-scope="scope">
          <el-button @click="edit(scope.row)" type="text" size="small">
            编辑
          </el-button>
          <el-button @click="delete_(scope.row)" type="text" size="small">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      background
      layout="prev, pager, next"
      :total="total"
      :page-size="pageSize"
      @current-change="page"
    >
    </el-pagination>
  </div>
</template>

<script>
export default {
  methods: {
    delete_(row) {
      this.$http.delete('http://localhost:8081/v1/dept/' + row.ID).then(resp => {
        console.log(resp)
        if (resp.data.message == 'ok') {
          this.$alert('删除成功', '系统提示', {
            confirmButtonText: '确定',
            callback: action => {
              window.location.reload();
            }
          })
        }
      })
    },
    edit(row) {
      this.$router.push({
        path: '/update',
        query: {
          ID: row.ID
        }
      })
    },
    page(currentPage) {
      this.$http.get('http://localhost:8081/v1/dept/' + currentPage + '/6').then(resp => {
        console.log(resp)
        this.tableData = resp.data.content
        this.pageSize = resp.data.size
        this.total = resp.data.totalElements
      })
    }
  },

  data() {
    return {
      pageSize: 1,
      total: 11,
      tableData: [{
        Deptno: '1',
        Dname: '研发',
        Loc: '北京',
      }, {
        Deptno: '2',
        Dname: '研发',
        Loc: '北京',
      }, {
        Deptno: '3',
        Dname: '研发',
        Loc: '北京',
      }]
    }
  },
  created() {
    this.$http.get('http://localhost:8081/v1/dept/1/6').then(resp => {
      console.log(resp)
      this.tableData = resp.data.content
      this.pageSize = resp.data.size
      this.total = resp.data.totalElements
    })
  }
}
</script>