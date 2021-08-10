<template>
  <div>
    <el-form
      :model="ruleForm"
      :rules="rules"
      ref="ruleForm"
      label-width="100px"
      class="demo-ruleForm"
    >
      <el-form-item label="部门名称" prop="Dname">
        <el-input v-model="ruleForm.Dname"></el-input>
      </el-form-item>
      <el-form-item label="地址" prop="Loc">
        <el-input v-model="ruleForm.Loc"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')"
          >立即创建</el-button
        >
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  data() {
    return {
      ruleForm: {
        Dname: '',
        Loc: '',
      },
      rules: {
        Dname: [
          { required: true, message: '请输入部门名称', trigger: 'blur' },
          { min: 0, max: 10, message: '长度在 0 到 10 个字符', trigger: 'blur' }
        ],
        Loc: [
          { required: true, message: '请输入部门地点', trigger: 'blur' }
        ],
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$http.post("http://localhost:8081/v1/dept", this.ruleForm).then(resp => {
            console.log(resp)
            if (resp.data.message == 'ok') {
              this.$alert('添加成功', '系统提示', {
                confirmButtonText: '确定',
                callback: action => {
                  this.$router.push('/Manager')
                }
              });
            }
          })
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
}
</script>