#include <iostream>
#include <random>
#include <string>
using namespace std;

class Card
{
	public:
		Card(int suit, int val) : m_suit(suit), m_val(val){}
		friend ostream& operator<<(ostream& , const Card&);
	private:
		int m_suit;
		int m_val;
};

class Player
{
	public:
		Player(){}
		~Player(){}
		void addCard(Card* new_card)
		{
			m_cards.push_back(new_card);
		}
		friend ostream& operator<<(ostream&, const Player&);
	private:
		std::vector<Card*> m_cards;
};

class Deck
{
	public:
		Deck()
		{
			for(int i =0; i < 52; i++)
			{
				Card card(i /13, i%13 );
				m_cards.push_back(card);
			}
		}
		void shuffle()
		{
			// Learn things from <random> 	
		}
	private:
		std::vector<Card> m_cards;
		int iter;
};

ostream& operator<<( ostream& os, const Card& card)
{
	std::string ranks[13] = {"A", "2", "3", "4", "5", "6",
							"7", "8", "9", "10", "J", "Q", "K"};
	std::string suits[4] = {"C", "D", "H", "S"};

	return os << ranks[card.m_val] << suits[card.m_suit];
}

ostream& operator<<( ostream& os, const Player& player)
{
	for(auto i : player.m_cards)
	{
		os << *i << " ";
	}
	return os;
}

std::vector<Player> players;
int player_count = 0;	

int main(int argc, char**argv){

	return 0;
}
