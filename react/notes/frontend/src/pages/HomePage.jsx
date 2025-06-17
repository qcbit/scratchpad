import React, { useState, useEffect } from 'react'
import { toast } from 'react-hot-toast';

import { Navbar } from '../components/Navbar';
import RateLimitedUI from '../components/RateLimitedUI';
import { NoteCard } from '../components/NoteCard';
import NotesNotFound from '../components/NotesNotFound';
import axiosInstance from '../lib/axios';

const HomePage = () => {
  const [isRateLimited, setIsRateLimited] = useState(false);
  const [notes, setNotes] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchNotes = async () => {
      try {
        const response = await axiosInstance.get("/notes");
        console.log(response.data);
        setNotes(response.data);
        setIsRateLimited(false);
      } catch (error) {
        console.error("Error fetching notes:", error);
        if (error.response?.status == 429) {
          setIsRateLimited(true);
        } else {
          toast.error("Failed to load notes. Please try again later.");
        }
      } finally {
        setLoading(false);
      }
    };
    fetchNotes();
}, []);

  return (
    <div className="min-h-screen">
      <Navbar />

      {isRateLimited && <RateLimitedUI />}

      <div className="max-w-7xl mx-auto p-4 mt-6">
        {loading && <div className="text-center text-primary py-10">Loading notes...</div>}

        {notes.length === 0 && !isRateLimited && <NotesNotFound />}

        {notes.length > 0 && !isRateLimited && (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {notes.map(note => (
              <div>
                <NoteCard key={note._id} note={note} setNotes={setNotes}/>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  )
};

export default HomePage;
