<!-- src/views/SignUp.vue -->
<template>
    <div class="signup">
      <h1>Sign Up</h1>
      <form @submit.prevent="submitForm">
        <div>
          <label for="username">Username</label>
          <input type="text" v-model="username" id="username" required />
        </div>
        <div>
          <label for="password">Password</label>
          <input type="password" v-model="password" id="password" required />
        </div>
        <button type="submit">Sign Up</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        username: '',
        password: '',
      };
    },
    methods: {
      async submitForm() {
        try {
          const response = await axios.post('http://localhost:8000/signup', {
            username: this.username,
            password: this.password,
          });
  
          if (response.status === 201) {
            alert('Sign Up successful');
            this.$router.push('/signin');
          }
        } catch (error) {
          console.error(error);
          if (error.response) {
            alert(`Sign Up failed: ${error.response.data}`);
          } else {
            alert('An error occurred during Sign Up');
          }
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .signup {
    max-width: 400px;
    margin: 50px auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 10px;
    background-color: #f9f9f9;
  }
  .signup form {
    display: flex;
    flex-direction: column;
  }
  .signup form div {
    margin-bottom: 15px;
  }
  .signup form label {
    margin-bottom: 5px;
    font-weight: bold;
  }
  .signup form input {
    padding: 8px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  .signup form button {
    padding: 10px;
    font-size: 16px;
    background-color: #333;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  .signup form button:hover {
    background-color: #555;
  }
  </style>
  