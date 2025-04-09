<script>
export default {
    data() {
        return {
            username: null,
            password: null,
        }
    },

    methods: {
        handleSubmit() {
            console.log(`Username: ${this.username}, Password: ${this.password}`);
            
        }    
    }
}

</script>

<template>
    <div class="login-page">
        <div class="container-auth-form">
            <div class="auth-form-content">
                <div class="head">
                    <div class="auth-form-login">
                        <div class="auth-form-title">
                            Login
                        </div>
                        <div class="auth-form-text">
                            Login and have more fun
                        </div>
                    </div>

                    <div class="button-back-to-home">
                        <div class="icon-auth-form">
                            <Img
                                alt="house-logo"
                                src="png/house.png"
                            />
                        </div>
                        <div class="button-back-to-home-text">
                           <a href="#"> Back to home </a>
                        </div>
                    </div>
                </div>

                <div class="inputs">
                    <Input
                        placeholder="Username"
                        v-model="username"
                    />

                    <Input
                        placeholder="Password"
                        v-model="password"
                    />
                </div>

                <div class="checkbox">
                    <div class="checkbox-click">
                        <input type="checkbox" id="remember-me" class="checkbox-input check-icon">
                        <label for="remember-me"></label>
                    </div>
                    
                    <div class="checkbox-text">
                        Remember me
                    </div>
                </div>

                
                
                <div class="button-login">
                        <Button
                            text="Login"
                            @submit="handleSubmit"
                        />    
                </div>

                <div class="register">
                    dont have an account?
                    <a href="#">Register</a> 
                </div>
            </div>

            <div class="img-login-page">
                <Img
                    alt="flowers"
                    src="image/imgLoginPage.png"
                />
            </div>
        </div>
    </div>
</template>


<style>

body {
    background-color: #F2F2F2;
    position: relative;
}

.login-page {
    display: flex;
    margin: 88px 232px;
    justify-content: space-between;
    background-color: #FFFFFF;
    max-height: 632px;
    max-width: 976px;
}

.container-auth-form {
    display: flex;
}

.icon-auth-form {
    max-height: 18px;
    display: flex;
    align-items: end;
    margin-right: 4px;
}

.inputs{
    display: flex;
    flex-direction: column;
    gap: 14px;
}

.button-back-to-home {
    display: flex;
    margin: 46px 30px 95px 40px;
}

.auth-form-login {
    margin-top: 33px;
}

.head {
    display: flex;
    justify-content: space-between;
}

.auth-form-title {
    font-size: 32px;
    line-height: 100%;
    letter-spacing: 0%;
    margin-bottom: 13px;
}

.auth-form-text {
    font-size: 18px;
    line-height: 100%;
    letter-spacing: 0%;
    color: #838383;

}

.checkbox-click {
    position: relative;
}

/* Чекбокс остается скрытым, но кликабельным */
.checkbox-click input {
    visibility: hidden; /* Скрыть чекбокс */
    position: absolute; /* Абсолютное позиционирование */
    width: 20px; /* Размер чекбокса */
    height: 20px; /* Высота чекбокса */
    z-index: 99; /* Проверяем, чтобы чекбокс находился сверху */
    cursor: pointer; /* Курсор мыши изменится на указатель */
}

/* Стили для SVG */
.checkbox-click label {
    display: inline-block;
    width: 20px;
    height: 20px;
    background-color: #D9D9D9;
    position: relative;
    cursor: pointer;
}

/* SVG для галочки */
.checkbox-click label::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='16' height='16' fill='none'%3E%3Cpath d='M12.7071 5.29289C13.0976 5.68342 13.0976 6.31658 12.7071 6.70711L7.70711 11.7071C7.31658 12.0976 6.68342 12.0976 6.29289 11.7071L3.29289 8.70711C2.90237 8.31658 2.90237 7.68342 3.29289 7.29289C3.68342 6.90237 4.31658 6.90237 4.70711 7.29289L7 9.58579L11.2929 5.29289C11.6834 4.90237 12.3166 4.90237 12.7071 5.29289Z' fill='white'/%3E%3C/svg%3E");
    background-size: contain;
    background-repeat: no-repeat;
    background-position: center;
    opacity: 0; /* Изначально невидимо */
}

/* Активируем чекбокс */
.checkbox-click input:checked + label {
    background-color: #FF8F52; /* Цвет фона чекбокса при активации */
}

/* Показываем галочку при активированном чекбоксе */
.checkbox-click input:checked + label::before {
    opacity: 1; /* Галочка становится видимой */
}

.checkbox {
    display: flex;
    margin-top: 23px;
    
}

.checkbox-text {
    font-size: 18px;
    line-height: 100%;
    letter-spacing: 0%;
    color: #838383;
    padding-left: 5px;
    display: flex;
    padding-top: 3px;
}

.button-login {
    margin-top: 66px;
    margin-bottom: 56px;
}

.register {
    font-size: 20px;
    color: #838383;
    padding: 0 40px 67px 57px;
}

.register a {
    color: #FF8F52;
}


.auth-form-content {
    padding: 0 48px 0 48px;
}






</style>
