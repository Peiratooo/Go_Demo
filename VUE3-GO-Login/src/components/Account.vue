<template>
    <div class="container">
        <div class="account">
            <div class="title">VU3 & Go</div>
            <router-view></router-view>
        </div>
    </div>
</template>

<script setup>

</script>

<style>
.container {
    height: 100vh;
    width: 100%;
    background-color: #2d2d4f;
    background-size: cover;
    background-position: center;
}

.account {
    display: flex;
    flex-direction: column;
    align-items: center;
    position: relative;
    top: 22vh;
}

.title {
    color: white;
    font-weight: bold;
    font-size: 26px;
    letter-spacing: 1px;
    margin-bottom: 25px;
}
.loginCard {
    display: flex;
    flex-direction: column;
    align-items: center;
    background-color: #20213d;
    padding: 36px 38px;
    border-radius: 15px;
    width: 20%;
    box-shadow:
        0px 0px 2.7px rgba(0, 0, 0, 0.031),
        0px 0px 6.9px rgba(0, 0, 0, 0.044),
        0px 0px 14.2px rgba(0, 0, 0, 0.056),
        0px 0px 29.2px rgba(0, 0, 0, 0.069),
        0px 0px 80px rgba(0, 0, 0, 0.1);
    margin-bottom: 26px;
}

span {
    cursor: pointer;
    text-decoration: underline;
}

.input {
    width: 80%;
    background-color: #454761;
    padding: 12px 18px;
    margin-bottom: 18px;
    border-radius: 8px;
}

.inputTitle {
    color: #b4b6c8;
    font-weight: bold;
    font-size: 12px;
    margin-bottom: 2px;
}

.input input {
    background-color: rgba(0, 0, 0, 0);
    width: 100%;
    border: 0px;
    border-bottom: 1px solid #2e2f4351;
    font-size: 16px;

}

.input input:focus {
    outline: none;
}

.input input::placeholder {
    color:#8889ab;
    font-size: 12px;
}

.button {
    background-color: #eff0ff;
    color: #11121f;
    font-weight: bold;
    padding: 10px 36px;
    border-radius: 30px;
    margin-top: 16px;
}
.subTitle,span {
    color: #7f7f97;
}
</style>