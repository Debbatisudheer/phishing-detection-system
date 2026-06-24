import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import api from "../services/api";

function SandboxReportDetails() {

  const { id } = useParams();

  const [report, setReport] =
    useState(null);

  useEffect(() => {

    loadReport();

  }, []);

  const loadReport =
    async () => {

      const response =
        await api.get(
          `/api/sandbox/report/${id}`,
        );

      setReport(
        response.data,
      );
    };

  if (!report) {

    return <h2>Loading...</h2>;
  }

  return (

    <div
      style={{
        padding: "20px",
      }}
    >

      <h1>
        Sandbox Report
      </h1>

      <p>
        <b>ID:</b> {report.id}
      </p>

      <p>
        <b>File:</b> {report.file_name}
      </p>

      <p>
        <b>Risk Score:</b> {report.risk_score}
      </p>

      <p>
        <b>Risk Level:</b> {report.risk_level}
      </p>

      <p>
        <b>Verdict:</b> {report.verdict}
      </p>

      <p>
        <b>MITRE:</b> {report.mitre}
      </p>

      <h3>
        Findings
      </h3>

      <pre>
        {report.findings}
      </pre>

    </div>
  );
}

export default SandboxReportDetails;