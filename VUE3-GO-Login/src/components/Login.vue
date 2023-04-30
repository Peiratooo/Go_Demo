<template>
    <div class="loginCard">
        <div class="input">
            <div class="inputTitle">账号/邮箱:</div>
            <input type="text" placeholder="请输入账号/邮箱" v-model="user.uidoremail">
        </div>
        <div class="input">
            <div class="inputTitle">密码:</div>
            <input type="password" placeholder="请输入密码" v-model="user.psw">
        </div>
        <div class="button" @click="login">登 录</div>
    </div>
    <div class="subTitle">没有账号？<span @click="router.push('/regist')">注 册</span></div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import {reactive} from 'vue'
const router = useRouter()
const user = reactive({
    uidoremail:"",
    psw:""
})
function login(){
    console.log(user);
}
</script>

<style>
</style>