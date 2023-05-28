/*
USER 
*/

//Регистрация нового пользователя:
fetch('/api/users', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      name: 'John Doe',
      email: 'john@example.com',
      password: 'password123',
      role: 'Candidate'
    }),
  })
  .then(response => response.json())
  .then(data => console.log('Success:', data))
  .catch((error) => console.error('Error:', error));


//Авторизация пользователя:
fetch('/api/users/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: 'john@example.com',
      password: 'password123'
    }),
  })
  .then(response => response.text())
  .then(data => {
    console.log('Success:', data);
    // Здесь вы можете сохранить полученный JWT токен, например, в localStorage:
    // localStorage.setItem('token', data);
  })
  .catch((error) => console.error('Error:', error));


//Получение информации о пользователе:
const userId = 1; // Замените на ID интересующего вас пользователя
fetch(`/api/users/${userId}`, {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
    // Если ваш бэкенд проверяет JWT токен, передайте его в заголовках:
    // 'Authorization': 'Bearer ' + localStorage.getItem('token')
  },
})
.then(response => response.json())
.then(data => console.log('Success:', data))
.catch((error) => console.error('Error:', error));


//Обновление роли пользователя:
const userId = 1; // Замените на ID интересующего вас пользователя
fetch(`/api/users/${userId}`, {
  method: 'PUT',
  headers: {
    'Content-Type': 'application/json',
    // Если ваш бэкенд проверяет JWT токен, передайте его в заголовках:
    // 'Authorization': 'Bearer ' + localStorage.getItem('token')
  },
  body: JSON.stringify({
    role: 'Intern'
  }),
})
.then(response => response.json())
.then(data => console.log('Success:', data))
.catch((error) => console.error('Error:', error));


//Обновление информации о пользователе:
const userId = 1; // Замените на ID интересующего вас пользователя
fetch(`/api/users/${userId}`, {
  method: 'PUT',
  headers: {
    'Content-Type': 'application/json',
    // Если ваш бэкенд проверяет JWT токен, передайте его в заголовках:
    // 'Authorization': 'Bearer ' + localStorage.getItem('token')
  },
  body: JSON.stringify({
    name: 'New Name',
    email: 'newemail@example.com',
    // и так далее
  }),
})
.then(response => response.json())
.then(data => console.log('Success:', data))
.catch((error) => console.error('Error:', error));


//Удаление пользователя:
const userId = 1; // Замените на ID интересующего вас пользователя
fetch(`/api/users/${userId}`, {
  method: 'DELETE',
  headers: {
    'Content-Type': 'application/json',
    // Если ваш бэкенд проверяет JWT токен, передайте его в заголовках:
    // 'Authorization': 'Bearer ' + localStorage.getItem('token')
  },
})
.then(response => response.json())
.then(data => console.log('Success:', data))
.catch((error) => console.error('Error:', error));


/*
Internship
*/