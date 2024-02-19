#include <iostream>
#include <string>
#include <vector>
#include <limits>

using namespace std;

enum class Category
{
  ELECTRONICS,
  APPAREL,
  FOOD,
  BOOKS,
  GROCERY,
  // Add more categories as needed
};

struct Item
{
  int id;
  string name;
  Category category;
  double price;
  int quantity;
};

vector<Item *> inventory; // Dynamic array of item pointers

int generateId()
{
  static int idCounter = 1;
  return idCounter++;
}

bool validateCategoryInput(int categoryInput)
{
  return categoryInput >= 0 && categoryInput <= static_cast<int>(Category::GROCERY);
}

bool validatePrice(double price)
{
  return price > 0.0;
}

bool validateQuantity(int quantity)
{
  return quantity > 0;
}

template <typename T>
T getValidatedInput(const string &prompt, bool (*validate)(T))
{
  T input;
  bool isValidInput = false;
  do
  {
    cout << prompt;
    cin >> input;
    if (cin.fail() || (cin.peek() != '\n' && cin.peek() != EOF))
    {
      cin.clear();
      cin.ignore(numeric_limits<streamsize>::max(), '\n');
      cerr << "Invalid input. Please try again." << endl;
    }
    else
    {
      if (validate(input))
      {
        isValidInput = true;
      }
      else
      {
        cerr << "Invalid input. Please try again." << endl;
        cin.clear();
        cin.ignore(numeric_limits<streamsize>::max(), '\n');
      }
    }
  } while (!isValidInput);
  return input;
}

void addItem(string name, Category category, double price, int quantity)
{
  try
  {
    Item *newItem = new Item{generateId(), name, category, price, quantity};
    inventory.push_back(newItem);
  }
  catch (const bad_alloc &e)
  {
    cerr << "Error allocating memory for new item: " << e.what() << endl;
  }
}

void addItemToInventory()
{
  string name;
  int categoryInput;
  Category category;
  double price;
  int quantity;

  cout << "Enter product name: ";
  cin >> name;

  categoryInput = getValidatedInput<int>("Enter category (0 for ELECTRONICS, 1 for APPAREL, 2 for FOOD, 3 for BOOKS, 4 for GROCERY): ", validateCategoryInput);
  category = static_cast<Category>(categoryInput);

  price = getValidatedInput<double>("Enter price: ", validatePrice);

  quantity = getValidatedInput<int>("Enter quantity: ", validateQuantity);

  addItem(name, category, price, quantity);
  cout << "Item added successfully." << endl;
}

void deleteItemById(int id)
{
  for (auto it = inventory.begin(); it != inventory.end(); ++it)
  {
    if ((*it)->id == id)
    {
      delete *it;
      inventory.erase(it);
      return;
    }
  }
  cout << "Item with ID " << id << " not found." << endl;
}

void viewInventory()
{
  for (const Item *item : inventory)
  {
    cout << "ID: " << item->id << ", Name: " << item->name
         << ", Category: " << static_cast<int>(item->category)
         << ", Price: " << item->price << ", Quantity: " << item->quantity << endl;
  }
}

void clearInventory()
{
  for (Item *item : inventory)
  {
    delete item;
  }
  inventory.clear();
}

int main()
{
  bool running = true;
  int choice;
  string name;
  Category category;
  double price;
  int quantity, id;

  while (running)
  {
    cout << "\nInventory Management System" << endl;
    cout << "1. Add Item" << endl;
    cout << "2. Delete Item" << endl;
    cout << "3. View Inventory" << endl;
    cout << "4. Exit" << endl;
    cout << "Enter your choice: ";
    cin >> choice;

    try
    {
      switch (choice)
      {
      case 1:
        // Add item
        addItemToInventory();
        break;
      case 2:
        // Delete item by ID
        cout << "Enter item ID to delete: ";
        cin >> id;
        deleteItemById(id);
        break;
      case 3:
        // View inventory
        viewInventory();
        break;
      case 4:
        // Exit program
        running = false;
        break;
      default:
        cout << "Invalid choice. Please try again." << endl;
      }
    }
    catch (const exception &e)
    {
      cerr << "An error occurred: " << e.what() << endl;
    }
  }

  // Clean up before exiting
  clearInventory();
  return 0;
}
