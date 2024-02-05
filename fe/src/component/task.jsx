import React, { useEffect, useState } from "react";
import {
  TextField,
  Button,
  List,
  ListItem,
  ListItemText,
  ListItemSecondaryAction,
  IconButton,
  Divider,
  ListItemButton,
  Collapse,
} from "@mui/material";
import ExpandLess from "@mui/icons-material/ExpandLess";
import ExpandMore from "@mui/icons-material/ExpandMore";
import DeleteIcon from "@mui/icons-material/Delete";
import FormDialog from "../component/dialog";
export default function TaskItem({
  task,
  onUpdateIsComplete,
  onDelete,
  onAddSubTask,
  onUpdate
}) {
  const [openDialog, setOpenDialog] = useState(false);
  const [newSubTask, setNewSubTask] = useState("");
  const [selectedIndex, setSelectedIndex] = useState(null);

  const handleAddSubTaskClick = () => {
    onAddSubTask(newSubTask);
    setNewSubTask("");
  };

  const handleUpdateClick = (index) => {
    
    setSelectedIndex(index);
    setOpenDialog(true);
  };

  const handleUpdateSubmit = (event,type) => {


    onUpdate(event,type)
  };
  const handleDelete = (type,id) => {
    onDelete(type,id)
  };

  const [open, setOpen] = useState(true);
  //
  const handleClick = () => {
    setOpen(!open);
  };

  return (
    <div style={{ marginTop: "10px" }}>
      <List>
        <ListItem>
          <ListItemText
            primary={`Task Name: ${task.taskName}`}
            secondary={task.deadline || "No Deadline"}
          />
          <ListItemSecondaryAction>
            <Button
              variant="contained"
              color="success" 
              style={{marginRight: "10px"}}
              onClick={() =>
                onUpdateIsComplete("task", task.taskId, task.isComplete)
              }
            >
              {task.isComplete ? "Mark Incomplete" : "Mark Complete"}
            </Button>
            {task.isComplete ? null : (
              <>
                <Button
                  variant="contained"
                  color="warning" 
                  style={{marginRight: "10px"}}
                  onClick={() => handleUpdateClick(task.taskId)}
                >
                  update
                </Button>
                {openDialog && selectedIndex === task.taskId && (
                  <FormDialog
                    open={openDialog}
                    setOpen={setOpenDialog}
                    handleSubmit={(v,type)=>{handleUpdateSubmit(v,type)}}
                    type="task"
                    valueData={task}
                    contentText="Update task"
                    submitButtonText="Save"
                  />
                )}
              </>
            )}

            <IconButton edge="end" aria-label="delete" onClick={()=>{handleDelete("task",task.taskId)}}>
              <DeleteIcon />
            </IconButton>
          </ListItemSecondaryAction>
        </ListItem>
      </List>

      <List style={{ paddingLeft: "20px" }}>
        <TextField
          label="Add Subtask"
          variant="outlined"
          value={newSubTask}
          onChange={(e) => setNewSubTask(e.target.value)}
          style={{ paddingRight: "20px" }}
        />
        <Button variant="contained" onClick={handleAddSubTaskClick}>
          Add Subtask
        </Button>
      </List>
      <List style={{ paddingLeft: "20px" }}>
        {task.subtasks && (
          <ListItemButton onClick={handleClick}>
            <ListItemText primary="subtasks" />
            {open ? <ExpandLess /> : <ExpandMore />}
          </ListItemButton>
        )}
        <Collapse in={open} timeout="auto" unmountOnExit>
          <div>
            {task.subtasks &&
              task.subtasks.map((subtask, index) => (
                <ListItem key={subtask.subtaskId}>
                  <ListItemText
                    style={{
                      textDecoration: subtask.isComplete
                        ? "line-through"
                        : "none",
                    }}
                    secondary={subtask.subTaskName}
                  />
                  <Button
                  color="success" 
                    style={{marginRight: "10px"}}
                    variant="contained"
                    onClick={() =>
                      onUpdateIsComplete(
                        "subtask",
                        subtask.subtaskId,
                        subtask.isComplete
                      )
                    }
                  >
                    {subtask.isComplete ? "Undo Complete" : "Complete"}
                  </Button>
                  <Button
                    color="warning" 
                    variant="contained"
                    onClick={() => handleUpdateClick(index)}
                  >
                    Update Subtask
                  </Button>
                  {openDialog && selectedIndex === index && (
                    <FormDialog
                      open={openDialog}
                      setOpen={setOpenDialog}
                      handleSubmit={(v,type)=>{handleUpdateSubmit(v,type)}}
                      type="subtask"
                      valueData={subtask}
                      contentText="Update Subtask"
                      submitButtonText="Save"
                    />
                  )}
                  <ListItemSecondaryAction>
                    <IconButton
                      edge="end"
                      aria-label="delete"
                      onClick={() => handleDelete("subtask", subtask.subtaskId)}
                    >
                      <DeleteIcon />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
              ))}
          </div>
        </Collapse>
      </List>
      <Divider style={{ margin: 10 }} />
    </div>
  );
}
