import java.util.Scanner;

public class CommunityNetwork {
	
	// instance variables
	private String name;
	private int age;
	private int zip;
	
	// constructor
	public CommunityNetwork(String n, int a, int z) {
		name = n;
		age = a;
		zip = z;
	}	
	
	public String getName() { 
		return name; 
	}
	
	public int getAge() { 
		return age; 
	}
	
	public int getZip() { 
		return zip; 
	}
	
	// method to determine if person can join network
	public static boolean eligibility(int age, int zip) {
		if (age >= 18 && zip > 10000) {
			return true;
		} else {
			return false;
		}
	}
	
	// main method
	public static void main(String[] args) {
		
		// declare Scanner
		Scanner scnr = new Scanner(System.in);
		
		// declare variables
		String name;
		int age;
		int zip;
		
		// prompt user for info
		System.out.print("Please enter your name: ");
		name = scnr.nextLine();
		
		System.out.print("Please enter your age: ");
		age = scnr.nextInt();
		
		System.out.print("Please enter your zip code: ");
		zip = scnr.nextInt();
		
		// create new CommunityNetwork object
		CommunityNetwork user = new CommunityNetwork(name, age, zip);
		
		// check if user is eligible
		if (eligibility(user.getAge(), user.getZip())) {
			System.out.println("You are eligible to join the Community Network!");
		} else {
			System.out.println("Sorry, you are ineligible to join the Community Network.");
		}	
		
		scnr.close();
	
	}
	
	// method to add user to network
	public static void addUser(String n, int a, int z) {
		// add user to network
		// write code for this
	}
	
	// method to remove user from network
	public static void removeUser(String n, int a, int z) {
		// remove user from network
		// write code for this
	}
	
	// method to check for duplicate user
	public static void checkDuplicate(String n, int a, int z) {
		// check for duplicate user
		// write code for this
	}
	
	// method to create groups
	public static void createGroup(String n) {
		// create group
		// write code for this
	}
	
	// method to delete a group
	public static void deleteGroup(String n) {
		// delete a group
		// write code for this
	}
	
	// method to search for user
	public static void searchUser(String n, int a, int z) {
		// search for user
		// write code for this
	}
	
	// method to search for group
	public static void searchGroup(String n) {
		// search for group
		// write code for this
	}

}