using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Backend.Persistence
{
    /// <inheritdoc />
    public partial class UpdateForeignkey : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Tasks_Priorities_FkPriorityId",
                table: "Tasks");

            migrationBuilder.DropForeignKey(
                name: "FK_Tasks_Types_FkTypeId",
                table: "Tasks");

            migrationBuilder.DropIndex(
                name: "IX_Tasks_FkPriorityId",
                table: "Tasks");

            migrationBuilder.DropIndex(
                name: "IX_Tasks_FkTypeId",
                table: "Tasks");

            migrationBuilder.DropColumn(
                name: "FkPriorityId",
                table: "Tasks");

            migrationBuilder.DropColumn(
                name: "FkTypeId",
                table: "Tasks");

            migrationBuilder.AddColumn<int>(
                name: "PriorityId",
                table: "Tasks",
                type: "INTEGER",
                nullable: false,
                defaultValue: 0);

            migrationBuilder.AddColumn<int>(
                name: "TypeId",
                table: "Tasks",
                type: "INTEGER",
                nullable: false,
                defaultValue: 0);

            migrationBuilder.UpdateData(
                table: "Tasks",
                keyColumn: "Id",
                keyValue: 1,
                columns: new[] { "DueDate", "PriorityId", "StartDate", "TypeId" },
                values: new object[] { new DateTime(2023, 12, 25, 7, 18, 58, 293, DateTimeKind.Utc).AddTicks(2125), 1, new DateTime(2023, 12, 25, 7, 18, 58, 293, DateTimeKind.Utc).AddTicks(2122), 1 });

            migrationBuilder.CreateIndex(
                name: "IX_Tasks_PriorityId",
                table: "Tasks",
                column: "PriorityId");

            migrationBuilder.CreateIndex(
                name: "IX_Tasks_TypeId",
                table: "Tasks",
                column: "TypeId");

            migrationBuilder.AddForeignKey(
                name: "FK_Tasks_Priorities_PriorityId",
                table: "Tasks",
                column: "PriorityId",
                principalTable: "Priorities",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_Tasks_Types_TypeId",
                table: "Tasks",
                column: "TypeId",
                principalTable: "Types",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Tasks_Priorities_PriorityId",
                table: "Tasks");

            migrationBuilder.DropForeignKey(
                name: "FK_Tasks_Types_TypeId",
                table: "Tasks");

            migrationBuilder.DropIndex(
                name: "IX_Tasks_PriorityId",
                table: "Tasks");

            migrationBuilder.DropIndex(
                name: "IX_Tasks_TypeId",
                table: "Tasks");

            migrationBuilder.DropColumn(
                name: "PriorityId",
                table: "Tasks");

            migrationBuilder.DropColumn(
                name: "TypeId",
                table: "Tasks");

            migrationBuilder.AddColumn<int>(
                name: "FkPriorityId",
                table: "Tasks",
                type: "INTEGER",
                nullable: true);

            migrationBuilder.AddColumn<int>(
                name: "FkTypeId",
                table: "Tasks",
                type: "INTEGER",
                nullable: true);

            migrationBuilder.UpdateData(
                table: "Tasks",
                keyColumn: "Id",
                keyValue: 1,
                columns: new[] { "DueDate", "FkPriorityId", "FkTypeId", "StartDate" },
                values: new object[] { new DateTime(2023, 11, 24, 16, 30, 11, 859, DateTimeKind.Local).AddTicks(6337), 1, 1, new DateTime(2023, 11, 24, 16, 30, 11, 859, DateTimeKind.Local).AddTicks(6324) });

            migrationBuilder.CreateIndex(
                name: "IX_Tasks_FkPriorityId",
                table: "Tasks",
                column: "FkPriorityId");

            migrationBuilder.CreateIndex(
                name: "IX_Tasks_FkTypeId",
                table: "Tasks",
                column: "FkTypeId");

            migrationBuilder.AddForeignKey(
                name: "FK_Tasks_Priorities_FkPriorityId",
                table: "Tasks",
                column: "FkPriorityId",
                principalTable: "Priorities",
                principalColumn: "Id");

            migrationBuilder.AddForeignKey(
                name: "FK_Tasks_Types_FkTypeId",
                table: "Tasks",
                column: "FkTypeId",
                principalTable: "Types",
                principalColumn: "Id");
        }
    }
}
