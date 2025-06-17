import Note from "../models/Note.js";

export async function getAllNotes(req, res) {
  try {
    const notes = await Note.find().sort({ createdAt: -1 }); // newest first
    res.status(200).json(notes);
  } catch (error) {
    console.error("Error fetching notes:", error);
    res.status(500).json({ message: "Internal server error" });
  }
};

export async function getNote(req, res) {
  try {
    const noteId = req.params.id;
    const note = await Note.findById(noteId);
    if (!note) {
      return res.status(404).json({ message: "Note "+noteId+" not found" });
    }
    res.status(200).json(note);
  } catch (error) {
    console.error("Error fetching note:", error);
    res.status(500).json({ message: "Internal server error" });
  }
}

export async function createNote(req, res) {
  try {
    const { title, content } = req.body;
    const note = new Note({ title, content });
    const savedNote = await note.save();
    res.status(201).json(savedNote);
  } catch (error) {
    console.error("Error creating note:", error);
    res.status(500).json({ message: "Internal server error" }); 
  }
};

export async function updateNote(req, res) {
  try {
    const { title, content } = req.body;
    const noteId = req.params.id;
    const updatedNote = await Note.findByIdAndUpdate(noteId, { title, content }, {new: true});
    if (!updatedNote) {
      return res.status(404).json({ message: "Note "+noteId+" not found" });
    }
    res.status(200).json(updatedNote);
  } catch (error) {
    console.error("Error updating note:", error);
    res.status(500).json({ message: "Internal server error" });
  }
};

export async function deleteNote(req, res) {
  try {
    const noteId = req.params.id;
    const deletedNote = await Note.findByIdAndDelete(noteId);
    if (!deletedNote) {
      return res.status(404).json({ message: "Note "+noteId+" not found" });
    }
    res.status(200).json({ message: "Note deleted successfully" });
  } catch (error) {
    console.error("Error deleting note:", error);
    res.status(500).json({ message: "Internal server error" });
  }
};
