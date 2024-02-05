import React, { useEffect, useState } from "react";
import {
  TextField,
  Button,
 
  Grid,
  Box,
  CardContent,
  Card,
  Typography,

} from "@mui/material";

import { deleteData, fetchingData, postData, putData } from "../utils/axios";
import TaskItem from "../component/task";


function TodoList() {
  const [nonCompleteTasks, setNonCompleteTasks] = useState([]);
  const [completeTasks, setCompleteTasks] = useState([]);
  const [updated, setUpdated] = useState(false);
  const [newTask, setNewTask] = useState({ taskName: "", deadline: "" });

  const handleChangeNewTask = (e) => {
    const { name, value } = e.target;
    setNewTask({
      ...newTask,
      [name]: value,
    });
   ;
  };

  const submitNewTask = async () => {
    try {
      await postData("task", newTask);
      setNewTask({ taskName: "", deadline: "" });
      setUpdated(true);
    } catch (error) {
      console.error("Error submitting new task: ", error);
    }
  };

  const fetchData = async () => {
    try {
      const data = await fetchingData("task");
      await setCompleteTasks(data.dto.completeTask);
      await setNonCompleteTasks(data.dto.nonCompleteTask);
    } catch (error) {
      console.error("Error fetching data: ", error);
    }
  };

  useEffect(() => {
    fetchData();
    console.log("Fetching data", completeTasks)
    setUpdated(false);
  }, [updated]);

  const handleDelete = async (type, id) => {
    try {
      await deleteData(`${type}/${id}`);
      setUpdated(true);
    } catch (error) {
      console.error("Error deleting task: ", error);
    }
  };

  const handleUpdateStatusIsComplete = async (type, id, statusData) => {
    try {
      const updateObj = {
        isComplete: !statusData,
        [`${type}Id`]: Number(id),
      };
      await putData(`${type}Complete`, updateObj);
      setUpdated(true);
    } catch (error) {
      console.error("Error updating status: ", error);
    }
  };
  const handleUpdateStatus = async (v, type) => {
    
    try {
      await putData(type, v);
      setUpdated(true);
    } catch (error) {
      console.error("Error updating status: ", error);
    }
  };

  const handleAddSubTask = async (taskId, subTaskName) => {
    try {
      await postData("subtask", { subTaskName, taskId });
      setUpdated(true);
    } catch (error) {
      console.error("Error adding subtask: ", error);
    }
  };

  return (
    <div style={{ margin: 20 }}>
      <Box>
        <h2>To-Do List</h2>
        <Grid container spacing={2}>
          <Grid item xs={3}>
            <TextField
              label="Task"
              variant="outlined"
              name="taskName"
              value={newTask.taskName}
              onChange={handleChangeNewTask}
            />
          </Grid>
          <Grid item xs={3}>
            <TextField
              label="Deadline"
              type="datetime-local"
              name="deadline"
              variant="outlined"
              value={newTask.deadline}
              onChange={handleChangeNewTask}
              InputLabelProps={{
                shrink: true,
              }}
            />
          </Grid>
          <Grid item xs={3}>
            <Button variant="contained" onClick={submitNewTask}>
              Add Task
            </Button>
          </Grid>
        </Grid>
        <Card style={{ marginTop: "25px" }}>
          <CardContent>
            <Typography variant="h5" component="h2">
              In Progress
            </Typography>
            {nonCompleteTasks && nonCompleteTasks.map((task) => (
              <TaskItem
                key={task.taskId}
                task={task}
                onUpdateIsComplete={handleUpdateStatusIsComplete}
                onDelete={(type,id) => handleDelete(type, id)}
                onAddSubTask={(subTaskName) =>
                  handleAddSubTask(task.taskId, subTaskName)
                }
                onUpdate={(v,type) => handleUpdateStatus(v, type)}
              />
            ))}
          </CardContent>
        </Card>
        <Card style={{ marginTop: "25px" }}>
          <CardContent>
            <Typography variant="h5" component="h2">
              Complete
            </Typography>
            {completeTasks && completeTasks.map((task,index) => (
              <TaskItem
                key={index}
                task={task}
                onUpdateIsComplete={handleUpdateStatusIsComplete}
                onDelete={(type,id) => handleDelete(type, id)}
                onAddSubTask={(subTaskName) =>
                  handleAddSubTask(task.taskId, subTaskName)
                }
              />
            ))}
          </CardContent>
        </Card>
      </Box>
    </div>
  );
}



export default TodoList;
