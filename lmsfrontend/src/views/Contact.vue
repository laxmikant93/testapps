<!-- src/views/Contact.vue -->
<template>
    <div class="contact">
      <h1>Contact Us</h1>
      <p>If you have any questions or feedback, feel free to reach out!</p>
      <form @submit.prevent="submitForm">
        <div>
          <label for="name">Name:</label>
          <input type="text" v-model="name" id="name" required />
        </div>
        <div>
          <label for="email">Email:</label>
          <input type="email" v-model="email" id="email" required />
        </div>
        <div>
          <label for="message">Message:</label>
          <textarea v-model="message" id="message" required></textarea>
        </div>
        <button type="submit">Send Message</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Contact',
    data() {
      return {
        name: '',
        email: '',
        message: '',
      };
    },
    methods: {
      submitForm() {
        // Handle form submission (e.g., send data to backend)
        alert(`Thank you, ${this.name}! Your message has been received.`);
        // Reset form fields
        this.name = '';
        this.email = '';
        this.message = '';
      },
    },
  };
  </script>
  
  <style scoped>
  .contact {
    max-width: 600px;
    margin: 50px auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 10px;
    background-color: #f9f9f9;
  }
  
  .contact h1 {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .contact form {
    display: flex;
    flex-direction: column;
  }
  
  .contact form div {
    margin-bottom: 15px;
  }
  
  .contact form label {
    margin-bottom: 5px;
    font-weight: bold;
  }
  
  .contact form input,
  .contact form textarea {
    padding: 8px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  .contact form button {
    padding: 10px;
    font-size: 16px;
    background-color: #333;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  
  .contact form button:hover {
    background-color: #555;
  }
  </style>
  