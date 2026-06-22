import { Link, useNavigate }
from "react-router-dom";

function Navbar() {

  const navigate =
    useNavigate();

  const logout = () => {

    localStorage.removeItem(
      "token",
    );

    navigate(
      "/login",
    );
  };

  return (

    <nav>

      <Link to="/">
        Dashboard
      </Link>

      {" | "}

      <Link to="/cases">
        Cases
      </Link>

      {" | "}

      <Link to="/search">
        Search
      </Link>

      {" | "}

      <Link to="/recent-findings">
        Recent Findings
      </Link>

      {" | "}

      <Link to="/update-case">
  Update Case
</Link>

{" | "}

<Link to="/export-iocs">
  Export IOC
</Link>

{" | "}

<Link to="/live-alerts">
  Live Alerts
</Link>

{" | "}

<Link to="/alert-history">
  Alert History
</Link>

{" | "}

<Link to="/threat-hunting">
  Threat Hunting
</Link>

{" | "}


<Link
  to="/api/threat-intel"
>
  Threat Intel
</Link>

{" | "}

<Link to="/mitre">
  MITRE
</Link>

{" | "}

<Link to="/mitre-heatmap">
  MITRE Heatmap
</Link>

{" | "}

<Link
  to="/incidents"
>
  Incidents
</Link>

{" | "}

<Link to="/correlation">
  IOC Correlation
</Link>

{" | "}

<Link to="/campaigns">
  Campaigns
</Link>

{" | "}

<Link to="/campaign-timeline">
  Campaign Timeline
</Link>

{" | "}

<Link to="/ioc-graph">
  IOC Graph
</Link>

{" | "}

<Link to="/ioc-trends">
  IOC Trends
</Link>

{" | "}

<Link to="/ioc-network">
  IOC Network
</Link>

{" | "}

<Link to="/investigation">
  Investigation
</Link>

{" | "}
      <button
        onClick={logout}
      >
        Logout
      </button>

    </nav>
  );
}

export default Navbar;