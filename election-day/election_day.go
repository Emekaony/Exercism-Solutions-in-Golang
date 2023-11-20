package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var p *int
	p = &initialVotes
	return p
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter != nil {
		return *counter
	} else {
		return 0
	}
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	// this is way easier, just an address!
	result := &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
	return result
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	name := result.Name
	votes := result.Votes
	message := fmt.Sprintf("%s (%d)", name, votes)
	return message
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	_, exists := results[candidate]
	if exists {
		// decrement the candidates votes by 1
		results[candidate]--
	}
}
