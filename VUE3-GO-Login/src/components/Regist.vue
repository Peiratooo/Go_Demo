<template>
    <div class="loginCard">
        <div class="input">
            <div class="inputTitle">邮箱:</div>
            <input type="text"  v-model="user.email">
        </div>
        <div class="input">
            <div class="inputTitle">密码:</div>
            <input type="password" v-model="user.psw">
        </div>
        <div class="button" @click="regist">注 册</div>
        
    </div>
    <div class="subTitle">已有账号？<span @click="router.push('/login')">登 录</span></div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import {reactive} from 'vue'
const router = useRouter()
const user = reactive({
    email:"",
    psw:""
})
function regist(){
    console.log(user);
}
</script>

<style>

</style>