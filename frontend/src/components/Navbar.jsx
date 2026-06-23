import { Link, useNavigate } from "react-router-dom";
import "./Navbar.css";

function Navbar() {

  const navigate =
    useNavigate();

  const logout = () => {

    localStorage.removeItem(
      "token"
    );

    navigate(
      "/login"
    );
  };

  return (

    <nav className="navbar">

      <Link
        to="/"
        className="nav-link"
      >
        Dashboard
      </Link>

      <div className="dropdown">

        <span className="nav-link">
          Detection ▼
        </span>

        <div className="dropdown-content">

          <Link to="/search">
            Search
          </Link>

          <Link to="/recent-findings">
            Recent Findings
          </Link>

          <Link to="/live-alerts">
            Live Alerts
          </Link>

          <Link to="/alert-history">
            Alert History
          </Link>

          <Link to="/update-case">
            Update Case
          </Link>

        </div>

      </div>

      <div className="dropdown">

        <span className="nav-link">
          Investigation ▼
        </span>

        <div className="dropdown-content">

          <Link to="/cases">
            Cases
          </Link>

          <Link to="/investigation">
            Investigation
          </Link>

          <Link to="/incidents">
            Incidents
          </Link>

          <Link to="/sandbox">
            Sandbox
          </Link>

        </div>

      </div>

      <div className="dropdown">

        <span className="nav-link">
          Threat Intel ▼
        </span>

        <div className="dropdown-content">

          <Link to="/api/threat-intel">
            Threat Intel
          </Link>

          <Link to="/correlation">
            IOC Correlation
          </Link>

          <Link to="/ioc-graph">
            IOC Graph
          </Link>

          <Link to="/ioc-trends">
            IOC Trends
          </Link>

          <Link to="/ioc-network">
            IOC Network
          </Link>

          <Link to="/export-iocs">
            Export IOC
          </Link>

        </div>

      </div>

      <div className="dropdown">

        <span className="nav-link">
          Campaigns ▼
        </span>

        <div className="dropdown-content">

          <Link to="/campaigns">
            Campaigns
          </Link>

          <Link to="/campaign-timeline">
            Campaign Timeline
          </Link>

        </div>

      </div>

      <div className="dropdown">

        <span className="nav-link">
          MITRE ▼
        </span>

        <div className="dropdown-content">

          <Link to="/mitre">
            MITRE Dashboard
          </Link>

          <Link to="/mitre-heatmap">
            MITRE Heatmap
          </Link>

        </div>

      </div>

      <Link
        to="/threat-hunting"
        className="nav-link"
      >
        Threat Hunting
      </Link>

      <button
        className="logout-btn"
        onClick={logout}
      >
        Logout
      </button>

    </nav>

  );
}

export default Navbar;