import axios from "axios";

const baseUrl = import.meta.env.VITE_API_URL;


const config = {
    headers: {
        'Content-Type': 'application/json',
      },
}

const fetchingData = async (link,queryParams) => {
  try {
    console.log("q :",queryParams)
    const response = await axios.get(`${baseUrl}/${link}`,{
      params: queryParams,
      headers: config.headers,
    })
    return response.data;
  } catch (error) {
    console.error(error);
    return error;
  }
};
const deleteData = async (link,params) => {
  try {
    const response = await axios.delete(`${baseUrl}/${link}`,params,config)
    return response.data;
  } catch (error) {
    console.error(error);
    return error;
  }
};

const postData = async (link,params) => {

  try {
      const resp = await axios.post(`${baseUrl}/${link}`, params,config)
      return resp.data;
  } catch (err) {
      // Handle Error Here
      console.log(err)
      return err;
  }
};

const putData = async (link,params) => {
  try {
    console.log("link",link)
      const resp = await axios.put(`${baseUrl}/${link}`, params,config) 
      return resp.data;
  } catch (err) {
      console.log(err)
      return err;
  }
};

export {postData,fetchingData,deleteData,putData};