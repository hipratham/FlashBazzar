// Placeholder for frontend JavaScript scripts
// Add your Alpine.js or vanilla JS code here to handle
// UI interactions, countdown timer, form submissions, etc.

// Example: Initialize Alpine.js (if using)
// import Alpine from 'alpinejs';
// window.Alpine = Alpine;
// Alpine.start();

// --- Firebase Initialization ---
// TODO: Add SDKs for Firebase products that you want to use
import { getAuth, createUserWithEmailAndPassword, signInWithEmailAndPassword, onAuthStateChanged } from "https://www.gstatic.com/firebasejs/11.8.1/firebase-auth.js";

// Initialize Firebase Auth
const auth = getAuth();

// Handle user authentication state changes
onAuthStateChanged(auth, (user) => {
    if (user) {
        // User is signed in
        console.log("User is signed in:", user);
        // Redirect to profile page or show user-specific content
        // window.location.href = '/profile.html'; // Example redirection
    } else {
        // User is signed out
        console.log("User is signed out");
        // Redirect to homepage or show sign-in form
        // window.location.href = '/'; // Example redirection
    }
});

// Handle Sign Up Form Submission
const signUpForm = document.getElementById('signup-form');
if (signUpForm) {
    signUpForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const email = signUpForm.email.value;
        const password = signUpForm.password.value;

        createUserWithEmailAndPassword(auth, email, password)
            .then((userCredential) => {
                // Signed up
                const user = userCredential.user;
                console.log("User signed up:", user);
                // You might want to save additional user info to Firestore here
            })
            .catch((error) => {
                const errorCode = error.code;
                const errorMessage = error.message;
                console.error("Sign up error:", errorCode, errorMessage);
                // Display error message to user
            });
    });
}

// Handle Sign In Form Submission
const signInForm = document.getElementById('signin-form');
if (signInForm) {
    signInForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const email = signInForm.email.value;
        const password = signInForm.password.value;

        signInWithEmailAndPassword(auth, email, password)
            .then((userCredential) => { /* Signed in */ console.log("User signed in:", userCredential.user); })
            .catch((error) => { console.error("Sign in error:", error.code, error.message); });
    });
}

// --- Deal Display Logic ---

// Placeholder function to fetch deal data (will be replaced with API call)
async function fetchDailyDeal() {
    // Simulate fetching data
    return new Promise(resolve => {
        setTimeout(() => {
            resolve({
                imageURL: 'https://via.placeholder.com/600x400?text=Daily+Deal+Image', // Replace with actual image URL
                title: 'Amazing Handcrafted Nepali Product!',
                originalPrice: 1500,
                flashPrice: 999,
                stock: 50,
                startTime: new Date().getTime() + 1000 * 60 * 5, // 5 minutes from now
                endTime: new Date().getTime() + 1000 * 60 * 60 * 24, // 24 hours from now
            });
        }, 500);
    });
}

// Function to display the deal data on the page
async function displayDailyDeal() {
    const deal = await fetchDailyDeal();

    if (!deal) {
        console.log("No deal available today.");
        // Display a message indicating no deal
        return;
    }

    document.getElementById('deal-image').src = deal.imageURL;
    document.getElementById('deal-title').textContent = deal.title;
    document.getElementById('original-price').textContent = `Rs. ${deal.originalPrice}`;
    document.getElementById('flash-price').textContent = `Rs. ${deal.flashPrice}`;
    document.getElementById('stock-badge').textContent = `Stock: ${deal.stock}`;

    // TODO: Implement countdown timer logic
    console.log("Deal start time:", new Date(deal.startTime));
}

displayDailyDeal();

// --- Buy Now Modal Logic ---

const buyNowButton = document.getElementById('buy-now-button');
const buyNowModal = document.getElementById('buy-now-modal');
const closeModalButton = document.getElementById('close-modal'); // Assuming you have a close button with this ID
const buyNowForm = document.getElementById('buy-now-form'); // Assuming the modal contains a form with this ID

// Function to show the modal
function showBuyNowModal() {
    buyNowModal.classList.remove('hidden');
}

// Function to hide the modal
function hideBuyNowModal() {
    buyNowModal.classList.add('hidden');
}

// Event listener for the "Buy Now" button
if (buyNowButton) {
    buyNowButton.addEventListener('click', showBuyNowModal);
}

// Event listener for the modal close button/area
if (closeModalButton) {
    closeModalButton.addEventListener('click', hideBuyNowModal);
    // Optionally add event listener to the modal overlay itself to close it
}

// --- Notification Subscription Logic ---

const notifyMeToggle = document.getElementById('notify-me-toggle'); // Assuming you have a checkbox or similar element with this ID

if (notifyMeToggle) {
    notifyMeToggle.addEventListener('change', (e) => {
        const isSubscribed = e.target.checked;
        console.log("Notification subscription toggled:", isSubscribed);

        // TODO: Send request to backend to update user's subscription status
        // You will need to send the user's UID and the new subscription status (isSubscribed)
        // Use fetch API or a library like Axios for the HTTP request
        // Handle success and error responses from the backend
    });
}

// Import the functions you need from the SDKs you need
import { initializeApp } from "https://www.gstatic.com/firebasejs/11.8.1/firebase-app.js";
import { getAnalytics } from "https://www.gstatic.com/firebasejs/11.8.1/firebase-analytics.js";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyBx_PgD8l6Afy10BqpnLsORaIqpypziuoQ",
  authDomain: "flashbazzar-761af.firebaseapp.com",
  projectId: "flashbazzar-761af",
  storageBucket: "flashbazzar-761af.firebasestorage.app",
  messagingSenderId: "700317985989",
  appId: "1:700317985989:web:da8e99ffeff75adeeb54ac",
  measurementId: "G-PZ0RBCGD0H"
};
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
