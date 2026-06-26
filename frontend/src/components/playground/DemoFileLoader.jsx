import api from "../../services/api";

async function loadDemoFile(path) {

  try {

    const response =
      await fetch(path);

    const blob =
      await response.blob();

    const file =
      new File(
        [blob],
        path.split("/").pop(),
      );

    const formData =
      new FormData();

    formData.append(
      "file",
      file,
    );

    const result =
      await api.post(
        "/api/analyze-file",
        formData,
        {
          headers: {
            "Content-Type":
              "multipart/form-data",
          },
        },
      );

    return result.data;

  } catch (error) {

    console.error(error);

    return null;

  }

}

export default loadDemoFile;