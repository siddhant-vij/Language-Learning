#include <cstdio>
#include <iostream>

int main()
{
  // Create and open a text file - write mode
  FILE *filePtr = fopen("lorem_ipsum.txt", "w");
  if (filePtr == nullptr)
  {
    perror("Error opening file");
    return 1;
  }

  // Write lorem ipsum text to the file
  const char *text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit.\n";
  fputs(text, filePtr);

  // Close the file
  fclose(filePtr);

  // Rename the file
  if (rename("lorem_ipsum.txt", "lorem_ipsum_renamed.txt") != 0)
  {
    perror("Error renaming file");
    return 1;
  }

  // Open the renamed file - read mode
  filePtr = fopen("lorem_ipsum_renamed.txt", "r");
  if (filePtr == nullptr)
  {
    perror("Error opening file");
    return 1;
  }

  // Read and display the content of the file
  char buffer[1024];
  while (fgets(buffer, 1024, filePtr) != nullptr)
  {
    std::cout << buffer;
  }

  // Close the file
  fclose(filePtr);

  // Open the file - append mode
  filePtr = fopen("lorem_ipsum_renamed.txt", "a");
  if (filePtr == nullptr)
  {
    perror("Error opening file");
    return 1;
  }

  // Append more lorem ipsum text to the file
  const char *moreText = "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.\n";
  fputs(moreText, filePtr);

  // Close the file
  fclose(filePtr);

  // Remove the file
  if (remove("lorem_ipsum_renamed.txt") != 0)
  {
    perror("Error deleting file");
    return 1;
  }

  std::cout << "File operations completed successfully." << std::endl;
  return 0;
}
