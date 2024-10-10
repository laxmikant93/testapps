<!-- src/components/Navbar.vue -->
<template>
    <nav class="navbar">
      <div class="container">
        <router-link to="/" class="logo">My Website</router-link>
        <ul class="nav-links">
          <li><router-link to="/">Home</router-link></li>
          <li><router-link to="/about">About</router-link></li>
          <li><router-link to="/services">Services</router-link></li>
          <li><router-link to="/contact">Contact</router-link></li>
          <li v-if="!isAuthenticated"><router-link to="/signin">Sign In</router-link></li>
          <li v-if="!isAuthenticated"><router-link to="/signup">Sign Up</router-link></li>
          <li v-if="isAuthenticated"><a href="#" @click.prevent="logout">Logout</a></li>
        </ul>
      </div>
    </nav>
  </template>
  
  <script>
  export default {
    name: 'Navbar',
    computed: {
      isAuthenticated() {
        return !!localStorage.getItem('token');
      },
    },
    methods: {
      logout() {
        localStorage.removeItem('token');
        this.$router.push('/');
      },
    },
  };
  </script>
  
  <style scoped>
  .navbar {
    background-color: #333;
    padding: 10px 0;
  }
  .container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
  }
  .logo {
    font-size: 24px;
    color: white;
    text-decoration: none;
  }
  .nav-links {
    list-style: none;
    display: flex;
    gap: 15px;
  }
  .nav-links a {
    color: white;
    text-decoration: none;
    padding: 8px 15px;
    border-radius: 5px;
    transition: background-color 0.3s ease;
  }
  .nav-links a:hover {
    background-color: #555;
  }
  </style>
  