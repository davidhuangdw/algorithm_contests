package google_code_jam.`2022`.`1c`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val IMP = "IMPOSSIBLE"
    val (T, ) = input_ints()
    (1..T).forEach {
        val cnt = hashMapOf<Char, Int>()
        val left = hashMapOf<Char, Char>()
        val right = hashMapOf<Char, Char>()
        val (N,) = input_ints()
        val S = split_input()
        fun set(): String {
            for(s in S){
                for(ch in s){
                    cnt[ch] = cnt.getOrDefault(ch, 0) + 1
                }
                for(i in 0 until s.length - 1){
                    val (a, b) = s[i] to s[i+1]
                    if(a == b) continue
                    if(a in right || b in left)
                        return IMP
                    right[a] = b
                    left[b] =a
                }
            }
            val done = hashSetOf<Char>()
            return buildString {
                while(true){
                    var st = '@'
                    for(ch in cnt.keys){
                        if(ch !in done && ch !in left){
                            st = ch
                            break
                        }
                    }
                    if(st == '@')
                        break
                    while(true){
                        done.add(st)
                        for(_i in 0 until cnt[st]!!)
                            append(st)
                        if(st !in right)
                            break
                        st = right[st]!!
                    }
                }
                if(done.size != cnt.size)
                    return IMP
            }
        }
        val res = set()
        println("Case #${it}: $res")
    }
}
