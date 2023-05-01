<template>
    <div class="home">
        <div class="homeContainer">
            <div class="items" v-for="(item, index) in data.results">
                <div class="item">{{ item.uid }}</div>
                <div class="item" style="margin-right: 60px;">{{ item.name }}</div>
                <div class="item">{{ item.email }}</div>
            </div>
        </div>
        <div class="exit" @click="exit()">
            退 出
        </div>
    </div>
</template>

<script setup>
import axios from 'axios';
import { onMounted, reactive } from "vue";
import { useRouter } from 'vue-router'
const router = useRouter()
const request = axios.create({
    baseURL: '/api/api',
    timeout: 1000,
    headers: {
        Authorization: localStorage.getItem('Token')
    }
})

const data = reactive({
    results: ""
})

function exit() {
    localStorage.removeItem('Token')
    router.push({
        name: "login",
    })
}

onMounted(() => {
    request.get('/user').then(res => {
        if (res.status == 200) {
            console.log(res.data.result);
            data.results = res.data.result
        } 
        if (res.data.msg == "TokenError") {
            exit()
        }
    })
})
</script>

<style scoped>
.home {
    background-color: #2d2d4f;
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    min-width: 580px;
    align-items: center;
    justify-content: center;
}

.homeContainer {
    display: flex;
    margin: 0 auto;
    flex-direction: column;
    align-items: center;
    background-color: #20213d;
    padding: 36px 38px;
    border-radius: 15px;
    width: 35%;
    min-width: 500px;
    margin-top: 100px;
    box-shadow:
        0px 0px 2.7px rgba(0, 0, 0, 0.031),
        0px 0px 6.9px rgba(0, 0, 0, 0.044),
        0px 0px 14.2px rgba(0, 0, 0, 0.056),
        0px 0px 29.2px rgba(0, 0, 0, 0.069),
        0px 0px 80px rgba(0, 0, 0, 0.1);
    margin-bottom: 26px;
}

.items {
    background-color: #454761;
    width: 80%;
    padding: 12px 20px;
    margin: 10px 0;
    border-radius: 6px;
    display: flex;
    align-items: center;
}

.item {
    margin-right: 10px;
    width: 80px;
}

.exit {
    margin-top: 20px;
    background-color: #454761;
    padding: 15px 38px;
    margin-bottom: 10vh;
    border-radius: 15px;
    cursor: pointer;
}</style>