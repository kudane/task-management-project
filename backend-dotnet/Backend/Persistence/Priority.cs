﻿using System.Reflection.Metadata;

namespace Backend.Persistence;

public partial class Priority
{
    public int Id { get; set; }
    public string? Name { get; set; }
    public virtual ICollection<Task> Tasks { get; set; } = new List<Task>();
}
