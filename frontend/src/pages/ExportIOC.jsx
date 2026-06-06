import api from "../services/api";

function ExportIOC() {

  const exportIOC = async () => {

    const token =
      localStorage.getItem(
        "token",
      );

    const response =
      await api.get(
        "/api/export/iocs",
        {
          headers: {
            Authorization:
              `Bearer ${token}`,
          },
          responseType:
            "blob",
        },
      );

    const url =
      window.URL.createObjectURL(
        new Blob(
          [response.data],
        ),
      );

    const link =
      document.createElement(
        "a",
      );

    link.href = url;

    link.download =
      "ioc_report.json";

    link.click();
  };

  return (
    <div>

      <h1>
        Export IOC Report
      </h1>

      <button
        onClick={exportIOC}
      >
        Download IOC Report
      </button>

    </div>
  );
}

export default ExportIOC;