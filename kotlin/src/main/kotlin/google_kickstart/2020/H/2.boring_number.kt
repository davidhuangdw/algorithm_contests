package google_kickstart.`2020`.H

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val Pow5 = LongArray(19)
    Pow5[0] = 1
    for(i in 1..18) Pow5[i] = Pow5[i-1] * 5

    fun boringCount(xx: Long): Long{
        if(xx == 0L) return 0L

        var x = xx
        val nums = mutableListOf<Long>()
        while(x > 0){
            nums.add(x % 10)
            x /= 10
        }
        nums.reverse()
        val n = nums.size
        var oddEven = 0
        while(oddEven < n && nums[oddEven] % 2 != oddEven % 2L) oddEven+=1

        var cnt = 0L
        for(zero in 1 until n) cnt += Pow5[n-zero]
        for(eq in 0..oddEven){
            cnt += if(eq < n){
                val needEven = if(eq %2 == 0) 0 else 1
                ((nums[eq]+needEven)/2) * Pow5[n-1-eq]
            }
            else 1
        }
        return cnt
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (L, R) = input_longs()
        println("Case #${it}: ${boringCount(R) - boringCount(L-1)}")
    }
}
