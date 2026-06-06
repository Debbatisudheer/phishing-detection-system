import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";

import Navbar from "./components/Navbar";

import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import Cases from "./pages/Cases";
import Search from "./pages/Search";
import RecentFindings from "./pages/RecentFindings";
import UpdateCase from "./pages/UpdateCase";

import ProtectedRoute from "./components/ProtectedRoute";
import ExportIOC
from "./pages/ExportIOC";
import LiveAlerts
from "./pages/LiveAlerts";
import ThreatHunting
from "./pages/ThreatHunting";
import FileDetails
from "./pages/FileDetails";
import ThreatIntel
from "./pages/ThreatIntel";
import MITREDashboard
from "./pages/MITREDashboard";
import IncidentDashboard
from "./pages/IncidentDashboard";

function App() {

  return (

    <BrowserRouter>

      <Navbar />

      <Routes>

        <Route
          path="/login"
          element={<Login />}
        />

        <Route
          path="/"
          element={
            <ProtectedRoute>
              <Dashboard />
            </ProtectedRoute>
          }
        />

        <Route
          path="/cases"
          element={
            <ProtectedRoute>
              <Cases />
            </ProtectedRoute>
          }
        />

        <Route
          path="/search"
          element={
            <ProtectedRoute>
              <Search />
            </ProtectedRoute>
          }
        />

        <Route
          path="/recent-findings"
          element={
            <ProtectedRoute>
              <RecentFindings />
            </ProtectedRoute>
          }
        />

        <Route
          path="/update-case"
          element={
            <ProtectedRoute>
              <UpdateCase />
            </ProtectedRoute>
          }
        />

        <Route
  path="/export-iocs"
  element={
    <ProtectedRoute>
      <ExportIOC />
    </ProtectedRoute>
  }
/>

<Route
  path="/live-alerts"
  element={
    <ProtectedRoute>
      <LiveAlerts />
    </ProtectedRoute>
  }
/>
<Route
  path="/threat-hunting"
  element={
    <ProtectedRoute>
      <ThreatHunting />
    </ProtectedRoute>
  }
/>

<Route
	path="/file/:fileName"
	element={
		<FileDetails />
	}
/>

<Route
  path="/threat-intel"
  element={
    <ThreatIntel />
  }
/>
<Route
  path="/mitre"
  element={
    <MITREDashboard />
  }
/>
<Route
  path="/incidents"
  element={
    <IncidentDashboard />
  }
/>

      </Routes>

    </BrowserRouter>
  );
}

export default App;