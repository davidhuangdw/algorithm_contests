package atcoder.a12

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
    repeat(inputInts()[0]){
        val (N) = inputInts()
        val A = inputInts().sorted()
        var sum = 0L
        var k = 0
        println( A.withIndex().maxOf{(i, v) ->
            sum += v
//            var (l, r) = 0 to i
//            while(l <= r){
//                val md = (l+r)/2
//                if(1L*(i+1)* A[md] <= sum){
//                    l = md +1
//                }else{
//                    r = md -1
//                }
//            }
//            i - r
            while(k < N && 1L * (i+1)*A[k] <= sum){
                k++
            }
            i+1 - k
        })
    }
}
