import React from "react";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom"; // Import Link
import Register from "../src/Authication/Login/Register/Register";
import Login from "../src/Authication/Login/Login/Login";
import RegisterPage from "../src/Authication/Login/pages/RegisterPage";
import Dashboard from "../src/Dashboard/Dashboard";
import "./App.css";

function App() {
  const styles = {
    container: {
      backgroundImage:
        'url("https://scontent.fcpt1-1.fna.fbcdn.net/v/t39.30808-6/302889114_381062770888851_8720314408090333185_n.jpg?_nc_cat=107&ccb=1-7&_nc_sid=efb6e6&_nc_ohc=wi9mUXHamuMAX_qBfoE&_nc_ht=scontent.fcpt1-1.fna&oh=00_AfBCf3o1UZBprwMReFWs1qSzvQCxZUrXgj3bMPsuJvdgCA&oe=65C08570")', // Set the path to your image
      backgroundSize: "cover", // Adjust the background size property as needed
      padding: "20px", // Add padding for better appearance
      backgroundColor: "black", // Set text color to be visible on the background
    },
    link: {
      color: "green", // Set the link color to green
      textDecoration: "underline", // Add underline for better visibility
    },
  };

  return (
    <div style={styles.container}>
      <Router>
        <div className="App">
          <Routes>
            <Route path="/Register" element={<Register />} />
            <Route path="/Login" element={<Login />} />
            <Route path="/RegisterPage" element={<RegisterPage />} />
            <Route path="/Dashboard" element={<Dashboard />} />
          </Routes>

          {/* Add Link tags */}
          <nav>
            <ul>
              <li>
                <Link to="/Register" style={styles.link}>
                  Register
                </Link>
              </li>
              <li>
                <Link to="/Login" style={styles.link}>
                  Login
                </Link>
              </li>
              <li>
                <Link to="/RegisterPage" style={styles.link}>
                  Register Page
                </Link>
              </li>
            </ul>
          </nav>
        </div>
      </Router>
    </div>
  );
}

export default App;
