/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l:=1
	r:=n
	for l<=r{
		m:=(r+l)/2
		
		if guess(m)==0 {
			return m
		} else if guess(m)==-1 {
			r=m-1
		} else {
			l=m+1
		}
	}	
	return l
}
