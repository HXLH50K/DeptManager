<template>
  <div>
    <table>
      <tr>
        <td>编号</td>
        <td>部门名称</td>
        <td>地址</td>
      </tr>
      <tr v-for="item in depts"  v-bind:key="item.id">
        <td>{{ item.Deptno }}</td>
        <td>{{ item.Dname }}</td>
        <td>{{ item.Loc }}</td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: 'Dept',
  data: function () {
    return {
      msg: '123',
      depts: [
        {
          deptno: 1,
          dname: '1yanfa',
          loc: '1beijing'
        }, {
          deptno: 2,
          dname: 'renshi',
          loc: 'beijing'
        }
      ]
    }
  },
  created: function () {
    console.log(123213)
    this.$http.get('http://localhost:8081/v1/dept').then(resp=>{
      console.log(resp);
      this.depts = resp.data;
    })
  }
}
</script>

<style>
</style>
