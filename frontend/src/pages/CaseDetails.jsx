import { useEffect, useState } from "react";
import api from "../services/api";

function CaseDetails() {

  const [caseData, setCaseData] =
    useState(null);

  useEffect(() => {

    const fetchCase =
      async () => {

        const token =
          localStorage.getItem(
            "token",
          );

        const response =
          await api.get(
            "/api/case-details/1",
            {
              headers: {
                Authorization:
                  `Bearer ${token}`,
              },
            },
          );

        setCaseData(
          response.data,
        );
      };

    fetchCase();

  }, []);

  if (!caseData)
    return <h2>Loading...</h2>;

  return (
    <div>

      <h1>
        Case Details
      </h1>

      <p>
        ID:
        {caseData.id}
      </p>

      <p>
        File:
        {caseData.file_name}
      </p>

      <p>
        Analyst:
        {caseData.analyst}
      </p>

      <p>
        Status:
        {caseData.status}
      </p>

      <p>
        Notes:
        {caseData.notes}
      </p>

      <p>
        Created:
        {caseData.created_at}
      </p>

    </div>
  );
}

export default CaseDetails;