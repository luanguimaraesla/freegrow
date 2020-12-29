const axios = require('axios')

const API_URL = process.env.VUE_APP_BACKEND_URL

exports.signin = (credentials) => {
  axios.post( `${API_URL}/signin`, credentials, {
    headers: { "content-type": "application/json"}
  }).then(result => {
      alert(JSON.stringify(result.data));
  }).catch(error => {
      console.error(error.response.data);
  });
}

exports.signup = (user) => {
  axios.post(`${API_URL}/signup`, user, {
    headers: { "content-type": "application/json" }
  }).then(result => {
      alert(JSON.stringify(result.data));
  }).catch(error => {
      console.error(error.response.data);
  });
}
