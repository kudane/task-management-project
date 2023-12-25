namespace Backend.Persistence;

public partial class Task
{
    public int Id { get; set; }
    public string? Name { get; set; }
    public string? Description { get; set; }
    public DateTime? StartDate { get; set; }
    public DateTime? DueDate { get; set; }
    public int TypeId { get; set; }
    public Type Type { get; set; } = null!;
    public int PriorityId { get; set; }
    public Priority Priority { get; set; } = null!;
}
