import axios from 'axios'

export const loginService = {
    login
};

async function login(email, password) {
    return await axios
    .post(process.env.VUE_APP_LOGIN_API_ENDPOINT, {
        email: email,
        password: password
    })
    .then(response => {
        if (response.status == 200) {
            return {
                success: true,
                token: response.data.token
            }
        } else {
            return {
                success: false,
                token: ""
            }
        }
    })
    .catch(error => {
        console.error(error);
        return {
            success: false,
            token: ""
        }
    })
}