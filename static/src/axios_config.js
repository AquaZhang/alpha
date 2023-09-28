import axios from "axios";

axios.defaults.baseURL = 'http://localhost:8000/api/v1/resource';
axios.defaults.timeout = 5000;
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
export default {
  dev: {
    '/': {
      target: 'http://localhost:8000/api/v1/resource',
      changeOrigin: true

    }
  }
}



