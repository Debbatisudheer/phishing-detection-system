import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";

import Navbar from "./components/Navbar";
import HealthBanner from "./components/common/HealthBanner";
import ProtectedRoute from "./components/ProtectedRoute";

// Pages
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import Cases from "./pages/Cases";
import Search from "./pages/Search";
import RecentFindings from "./pages/RecentFindings";
import UpdateCase from "./pages/UpdateCase";
import ExportIOC from "./pages/ExportIOC";
import LiveAlerts from "./pages/LiveAlerts";
import ThreatHunting from "./pages/ThreatHunting";
import FileDetails from "./pages/FileDetails";
import ThreatIntel from "./pages/ThreatIntel";
import MITREDashboard from "./pages/MITREDashboard";
import IncidentDashboard from "./pages/IncidentDashboard";
import AlertHistory from "./pages/AlertHistory";
import Correlation from "./pages/Correlation";
import Campaigns from "./pages/Campaigns";
import IOCGraph from "./pages/IOCGraph";
import MITREHeatmap from "./pages/MITREHeatmap";
import CampaignTimeline from "./pages/CampaignTimeline";
import IOCTrends from "./pages/IOCTrends";
import IOCNetworkGraph from "./pages/IOCNetworkGraph";
import InvestigationWorkbench from "./pages/InvestigationWorkbench";
import SandboxDashboard from "./pages/SandboxDashboard";
import SandboxReportDetails from "./pages/SandboxReportDetails";
import Playground from "./pages/Playground";
import NotFound from "./pages/NotFound";

function App() {
  return (
    <BrowserRouter>
      {/* Global Application UI */}
      <HealthBanner />
      <Navbar />

      <Routes>
        {/* ============================================================
            Authentication
            ============================================================ */}
        <Route
          path="/login"
          element={<Login />}
        />

        {/* ============================================================
            Protected Application Routes
            ============================================================ */}
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
          path="/playground"
          element={
            <ProtectedRoute>
              <Playground />
            </ProtectedRoute>
          }
        />

        {/* ============================================================
            File Investigation
            ============================================================ */}
        <Route
          path="/file/:fileName"
          element={<FileDetails />}
        />

        {/* ============================================================
            Threat Intelligence
            ============================================================ */}
        <Route
          path="/api/threat-intel"
          element={<ThreatIntel />}
        />

        {/* ============================================================
            MITRE ATT&CK
            ============================================================ */}
        <Route
          path="/mitre"
          element={<MITREDashboard />}
        />

        <Route
          path="/mitre-heatmap"
          element={<MITREHeatmap />}
        />

        {/* ============================================================
            Incident Management
            ============================================================ */}
        <Route
          path="/incidents"
          element={<IncidentDashboard />}
        />

        <Route
          path="/alert-history"
          element={<AlertHistory />}
        />

        <Route
          path="/correlation"
          element={<Correlation />}
        />

        {/* ============================================================
            Campaign Analysis
            ============================================================ */}
        <Route
          path="/campaigns"
          element={<Campaigns />}
        />

        <Route
          path="/campaign-timeline"
          element={<CampaignTimeline />}
        />

        {/* ============================================================
            IOC Analysis & Visualization
            ============================================================ */}
        <Route
          path="/ioc-graph"
          element={<IOCGraph />}
        />

        <Route
          path="/ioc-trends"
          element={<IOCTrends />}
        />

        <Route
          path="/ioc-network"
          element={<IOCNetworkGraph />}
        />

        {/* ============================================================
            Investigation
            ============================================================ */}
        <Route
          path="/investigation"
          element={<InvestigationWorkbench />}
        />

        {/* ============================================================
            Sandbox
            ============================================================ */}
        <Route
          path="/sandbox"
          element={<SandboxDashboard />}
        />

        <Route
          path="/sandbox-report/:id"
          element={<SandboxReportDetails />}
        />

        {/* ============================================================
            Fallback
            ============================================================ */}
        <Route
          path="*"
          element={<NotFound />}
        />
      </Routes>
    </BrowserRouter>
  );
}

export default App;