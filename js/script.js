const loginBtn = document.getElementById('loginBtn');
const signupBtn = document.getElementById('signupBtn')

loginBtn.addEventListener('click', function() {
    document.querySelector('.modal').style.display = 'flex';
})

signupBtn.addEventListener('click', function() {
    document.querySelector('.modal2').style.display = 'flex';
})

document.querySelector('.close').addEventListener('click', function() {
    document.querySelector('.modal').style.display = 'none';
})

document.querySelector('.close2').addEventListener('click', function() {
    document.querySelector('.modal2').style.display = 'none';
})
